package user

import (
	"cyc/goblog/pkg/logger"
	"cyc/goblog/pkg/model"
	"cyc/goblog/pkg/types"
)

func Get(idstr string) (User, error)  {
	var userStruct User

	id := types.StringToInt(idstr)

	if err := model.DB.First(&userStruct, id).Error; err != nil {
		return userStruct, err
	}
	return userStruct, nil
}

func GetByEmail(email string) (User, error)  {
	var userStruct User

	if err := model.DB.Where("email", email).First(&userStruct).Error; err != nil {
		return userStruct, err
	}

	return userStruct, nil
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error)  {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}