package user

import (
	"cyc/goblog/app/models"
	password2 "cyc/goblog/pkg/password"
)

type User struct {
	models.BaseModel


	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (user User) ComparePassword(password string) bool  {
	return password2.CheckHash(password, user.Password)
}
