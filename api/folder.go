package api

import (
	"github.com/gin-gonic/gin"
	"myfile/service"
)

// FolderCreate 创建文件夹
func FolderCreate(c *gin.Context) {
	var service service.FolderCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateFolder()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// FolderUpdate 修改文件夹
func FolderUpdate(c *gin.Context) {
	var service service.FolderUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateFolder(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// FolderList 获取文件夹列表
func FolderList(c *gin.Context) {
	var service service.FolderListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListFolder(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// FolderInfo 获取文件夹详情
func FolderInfo(c *gin.Context) {
	var service service.FolderInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InfoFolder(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
