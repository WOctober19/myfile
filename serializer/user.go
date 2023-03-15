package serializer

import (
	"myfile/model"
	"strings"
)

// User 用户序列化器
type User struct {
	ID        uint     `json:"id"`
	UserName  string   `json:"user_name"`
	Nickname  string   `json:"nickname"`
	Status    string   `json:"status"`
	Avatar    string   `json:"avatar"`
	Token     string   `json:"token"`
	Access    []string `json:"access"`
	CreatedAt int64    `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	var access = strings.Split(user.Access, ",")
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Token:     user.Token,
		Access:    access,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
