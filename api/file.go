package api

import (
	"github.com/gin-gonic/gin"
)

// FileUpload 文件上传
func FileUpload(c *gin.Context) {
	var service file.FileUploadService
	if err := c.ShouldBind(&service); err == nil {
		service.UploadFile(c)
		//c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// FileList 文件列表
func FileList(c *gin.Context) {
	var service file.FileListService
	if err := c.ShouldBind(&service); err == nil {
		service.FileList(c)
		//c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// FileList 文件列表
func FileInfo(c *gin.Context) {
	var service file.FileInfoService
	if err := c.ShouldBind(&service); err == nil {
		service.FileInfo(c.Param("id"))
		//c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
