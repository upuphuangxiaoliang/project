package dto

import (
	"project/project2/model"
)

// 这里进行一个转换 不然用户的信息全部暴露出去
type UserDto struct {
	Name      string
	Telephone string
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
