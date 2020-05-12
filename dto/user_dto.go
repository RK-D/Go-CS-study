package dto

import "pers.study/cstest/model"

//数据转换，返回给前端只给那部分前端需要的，对其他用户信息进行加密保护
type UserDTO struct {
	Name string `json:"name"`
	//grom:不要加空格
	Telephone string `json:"telephone"`
}

func ConvertToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
