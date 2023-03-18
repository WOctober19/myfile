package model

import (
	"fmt"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Uid      int64  `gorm:"comment:用户ID"`
	Fid      int64  `gorm:"comment:文件夹ID"`
	Path     string `gorm:"comment:文件路径"`
	FileName string `gorm:"comment:文件名"`
	SaveName string `gorm:"comment:保存文件名"`
	Ext      string `gorm:"comment:文件类型"`
	Size     int64  `gorm:"comment:文件大小"`
	Sha1     string `gorm:"comment:文件hash"`
}

func FastUploadFile(sha1 string) (bool, File) {
	var file File
	res := DB.First(&file, "sha1 = ?", sha1)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	if res.RowsAffected == 0 {
		return false, file
	}
	return true, file
}
