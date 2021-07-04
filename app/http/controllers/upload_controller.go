package controllers

import (
	"cyc/goblog/pkg/config"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type UploadController struct {
	BaseController
}

type UploadResData struct {
	Code int
	Msg string
	Data map[string]interface{}
}

func (u *UploadController) UploadImage(w http.ResponseWriter, r *http.Request)  {
	res := UploadResData{
		Code:1,
		Msg: "success",
		Data: map[string]interface{}{
			"success": 1,
			"message": "成功",
			"url": "",
		},
	}

	w.Header().Set("Content-type", "text/html")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		// 接收图片
		uploadFile, handle, err := r.FormFile("editormd-image-file")

		if err != nil {
			res.Data["success"] = 0
			res.Data["message"] = "图片上传失败"
			data, _ := json.Marshal(res)
			w.Write(data)
		} else {
			// 图片逻辑处理
			ext := strings.ToLower(path.Ext(handle.Filename))
			if ext != ".jpg" && ext != ".jpeg" && ext != ".gif" && ext != ".png" {
				res.Data["success"] = 0
				res.Data["message"] = "图片格式不正确"
				data, _ := json.Marshal(res)
				w.Write(data)
				return
			}

			// 保存图片
			filePath := "./public/editor_md/images/uploads/"
			fileUrl := filePath + handle.Filename

			err = os.Mkdir(filePath, 0777)

			saveFile, err := os.OpenFile(fileUrl, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				res.Data["success"] = 0
				res.Data["message"] = "文件打开失败"
				data, _ := json.Marshal(res)
				w.Write(data)
				return
			}

			// 写入内容
			_, err = io.Copy(saveFile, uploadFile)
			if err != nil {
				res.Data["success"] = 0
				res.Data["message"] = "文件写入失败"
				data, _ := json.Marshal(res)
				w.Write(data)
				return
			}

			res.Data["url"] = config.GetString("app.url") + "/" + strings.TrimLeft(fileUrl, "./public")

			defer uploadFile.Close()
			defer saveFile.Close()

		}

	}
	res.Data["success"]= 1
	res.Data["message"] = "上传文件成功"

	data, _ := json.Marshal(res.Data)

	w.Write(data)
}