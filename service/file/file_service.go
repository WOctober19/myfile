package service

import (
	"github.com/gin-gonic/gin"
	"myfile/model"
	"myfile/serializer"
)

// FileListService 接收参数设置
type FileListService struct {
}

type FileInfoService struct {
}

// FileList 文件列表
func (service *FileListService) FileList(c *gin.Context) serializer.Response {
	var fileList []model.File
	fid, ok := c.GetQuery("folder_id")
	if !ok {
		return serializer.Response{
			Code: 500,
			Msg:  "未接收到参数!",
		}
	}

	err := model.DB.Where("uid = ?", fid).Find(&fileList).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "查询失败!",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFileListResponse(fileList),
	}
}

// FileInfo 文件详情
func (service *FileInfoService) FileInfo(id string) serializer.Response {
	fileObj := model.File{}

	if err := model.DB.First(&fileObj, id).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "文件不存在!",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: fileObj,
	}
}
