package model

import (
	"fmt"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Uid      int64
	Path     string
	FileName string
	SaveName string
	Ext      string
	Size     int64
	Sha1     string
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
