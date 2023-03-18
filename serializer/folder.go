package serializer

import "myfile/model"

// Folder 文件夹序列化器
type Folder struct {
	ID        uint   `json:"id"`
	Uid       int64  `json:"uid"`
	Introduce string `json:"introduce"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Size      int64  `json:"size"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"created_at"`
}

func BuildFolder(folder model.Folder) Folder {
	return Folder{
		ID:        folder.ID,
		Uid:       folder.Uid,
		Introduce: folder.Introduce,
		Name:      folder.Name,
		Password:  folder.Password,
		Size:      folder.Size,
		CreatedAt: folder.CreatedAt.Unix(),
		UpdatedAt: folder.UpdatedAt.Unix(),
	}
}

// BuildFolderResponse 序列化文件夹响应
func BuildFolderResponse(folder model.Folder) Response {
	return Response{
		Data: BuildFolder(folder),
	}
}

// BuildFolderListResponse 序列化文件夹列表响应
func BuildFolderListResponse(folderList []model.Folder) (list []Folder) {
	for _, item := range folderList {
		folder := BuildFolder(item)
		list = append(list, folder)
	}
	return list
}
