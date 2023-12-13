package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"userName"`
	PasswordHash string `json:"passwordHash"`
}

// 根据用户名判断用户是否存在，如果存在就返回1,否则返回0，如果返回-1，就是函数执行出错
func IsExist(userName string) (int, *User) {
	u := User{}
	result := Db.Table("users").Where("user_name = ?", userName).Find(&u)
	if result.Error != nil {
		return -1, nil
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return 1, &u
}

func Add(u *User) error {
	err := Db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
