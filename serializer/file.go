package serializer

import "myfile/model"

// File 文件序列化器
type File struct {
	ID        uint   `json:"id"`
	Path      string `json:"path"`
	FileName  string `json:"filename"`
	SaveName  string `json:"savename"`
	Ext       string `json:"ext"`
	Size      int64  `json:"size"`
	Sha1      string `json:"ext"`
	CreatedAt int64  `json:"created_at"`
}

// BuildFile 序列化文件
func BuildFile(file model.File) File {
	return File{
		ID:        file.ID,
		Path:      file.Path,
		FileName:  file.FileName,
		SaveName:  file.SaveName,
		Ext:       file.Ext,
		Size:      file.Size,
		Sha1:      file.Sha1,
		CreatedAt: file.CreatedAt.Unix(),
	}
}

// BuildFileListResponse 序列化文件列表
func BuildFileListResponse(folderList []model.File) (list []File) {
	for _, item := range folderList {
		folder := BuildFile(item)
		list = append(list, folder)
	}
	return list
}
