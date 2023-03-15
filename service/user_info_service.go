package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myfile/model"
	"myfile/serializer"
)

// UserLoginService 管理用户登录的服务
type UserInfoService struct {
	Token string `form:"token" json:"token" binding:"required"`
}

// setSession 设置session
func (service *UserInfoService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// UserInfo 用户信息函数
func (service *UserInfoService) UserInfo(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("token = ?", service.Token).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
