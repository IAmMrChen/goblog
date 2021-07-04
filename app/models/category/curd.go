package category

import (
	"cyc/goblog/pkg/logger"
	"cyc/goblog/pkg/model"
	"cyc/goblog/pkg/types"
)

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (category *Category) Create() (err error) {
	if err := model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// All 获取分类数据
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// 切片map的形式返回数据
func AllForSliceMap() ([]map[string]interface{}, error)  {
	resultData := []map[string]interface{}{}

	categoryData, err := All()
	if err != nil {
		return resultData, err
	}
	for _, value := range categoryData {
		tempMap := map[string]interface{}{}

		tempMap["id"] = value.ID
		tempMap["name"] = value.Name
		tempMap["is_choose"] = false // 用来判断这挑数据是否被选中，体现在编辑时候默认选中
		tempMap["created_at"] = value.CreatedAt

		resultData = append(resultData, tempMap)
	}

	return resultData, nil
}

// Get 通过 ID 获取分类
func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToInt(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}
