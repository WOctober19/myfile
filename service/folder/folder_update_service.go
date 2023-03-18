package service

import (
	"myfile/model"
	"myfile/serializer"
	"myfile/util"
)

// FolderUpdateService 接收参数设置
type FolderUpdateService struct {
	Name      string `form:"name" json:"name"`
	Introduce string `form:"introduce" json:"introduce"`
	Password  string `form:"password" json:"password"`
}

// UpdateFolder 创建文件夹
func (service *FolderUpdateService) UpdateFolder(id string) serializer.Response {
	folderObj := model.Folder{}

	if err := model.DB.First(&folderObj, id).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "文件夹不存在!",
			Error: err.Error(),
		}
	}

	folderObj.Name = service.Name
	folderObj.Introduce = service.Introduce
	if service.Password != "" {
		folderObj.Password = util.Md5Str(service.Password)
	}

	if err := model.DB.Save(&folderObj).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "文件夹更新失败!",
			Error: err.Error(),
		}
	}

	return serializer.BuildFolderResponse(folderObj)
}
