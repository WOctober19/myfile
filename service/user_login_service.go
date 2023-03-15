package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"myfile/model"
	"myfile/serializer"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var nowDate = time.Now().Format("2006-01-02 15")
var secret = fmt.Sprintf("%v%v", nowDate, "3333")

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"userName" json:"userName" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	dict := make(map[string]interface{})
	dict["name"] = user.UserName
	dict["password"] = user.PasswordDigest

	user.Token, _ = GenerateToken(dict, secret)
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.ParamErr("代码错误", nil)
	}
	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}

// token: "eyJhbGciO...解析token"
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["cmd"].(string), nil
}
