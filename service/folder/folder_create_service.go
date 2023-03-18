package service

import (
	"myfile/model"
	"myfile/serializer"
)

// FolderCreateService 接收参数设置
type FolderCreateService struct {
	Uid       int64  `form:"uid" json:"uid" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	Introduce string `form:"introduce" json:"introduce" binding:"required"`
}

// CreateFolder 创建文件夹
func (service *FolderCreateService) CreateFolder() serializer.Response {
	folderObj := model.Folder{}

	folderObj.Name = service.Name
	folderObj.Introduce = service.Introduce
	folderObj.Uid = service.Uid
	folderObj.Size = 0

	if err := model.DB.Create(&folderObj).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "创建文件夹失败!",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建文件夹成功!",
	}
}
