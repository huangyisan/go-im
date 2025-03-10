/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import (
	"encoding/json"
	"go_im/pkg/model"
)

type Users struct {
	ID              uint64 `json:"id"`
	Email           string `valid:"email" json:"email"`
	Password        string `valid:"password"`
	Avatar          string `json:"avatar"`
	Name            string `json:"name" valid:"name"`
	OauthType       int
	Status          int `json:"status"`
	OauthId         string
	CreatedAt       string `json:"created_at"`
	PasswordComfirm string ` gorm:"-" valid:"password_comfirm"`
	Bio string `json:"bio"`
	Sex int `json:"sex"`
	ClientType int `json:"client_type"`
	Age int `json:"age"`
}

type UsersWhiteList struct {
	ID              uint64 `json:"id"`
	Email           string `valid:"email" json:"email"`
	Avatar          string `json:"avatar"`
	Name            string `json:"name" valid:"name"`
	OauthType       int
	Status          int `json:"status"`
	OauthId         string
	CreatedAt       string `json:"created_at"`
	Bio string `json:"bio"`
	Sex int `json:"sex"`
	ClientType int `json:"client_type"`
	Age int `json:"Age"`
}
// 字段过滤机制
func (u *Users) MarshalJSON() ([]byte, error) {
	// 将 User 的数据映射到 UsersWhiteList 上
	user := UsersWhiteList{
		ID:       u.ID,
		Email: u.Email,
		Name: u.Name,
		Avatar:     u.Avatar,
		CreatedAt:     u.CreatedAt,
		Bio:     u.Bio,
		Sex:     u.Sex,
		ClientType:     u.ClientType,
		Status:     u.Status,
		Age:     u.Age,
	}
	return json.Marshal(user)
}

func (Users) TableName() string {
	return "im_users"
}


// 当前登录用户
var AuthUser *Users

func (a Users) GetAvatar() string {
	if a.Avatar == "" {
		return "https://learnku.com/users/27407"
	}
	return a.Avatar
}

func GetFriendListV2(user_id[] uint64)  ([]Users,error) {
	var users []Users
	err := model.DB.Where("id in (?)",user_id).Find(&users).Error;
	if err!=nil{
		return users,err
	}
	return users,nil
}

//设置用户上下线状态
func SetUserStatus(id uint64 ,status int )  {
	model.DB.Model(&Users{}).Where("id=?",id).Update("status",status)
}
