package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// UserRegister 用户注册接口
func FileUpload(c *gin.Context) {
	var service service.FileUploadService
	if err := c.ShouldBind(&service); err == nil {
		service.UploadFile(c)
		//c.JSON(200, res)
	} else {
		//c.JSON(200, ErrorResponse(err))
	}
}
