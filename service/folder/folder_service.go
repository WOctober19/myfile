package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myfile/model"
	"myfile/serializer"
)

// FolderListService 接收参数设置
type FolderListService struct {
}

type FolderInfoService struct {
}

// ListFolder 文件夹列表
func (service *FolderListService) ListFolder(c *gin.Context) serializer.Response {
	var folderList []model.Folder
	uid, ok := c.GetQuery("uid")
	if !ok {
		return serializer.Response{
			Code: 500,
			Msg:  "未接收到参数!",
		}
	}
	fmt.Println(uid)
	err := model.DB.Where("uid = ?", uid).Find(&folderList).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "查询失败!",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFolderListResponse(folderList),
	}
}

// InfoFolder 文件夹详情
func (service *FolderInfoService) InfoFolder(id string) serializer.Response {
	folderObj := model.Folder{}

	if err := model.DB.First(&folderObj, id).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "文件夹不存在!",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: folderObj,
	}
}
