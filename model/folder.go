package model

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Uid       int64  `gorm:"comment:用户ID"`
	Introduce string `gorm:"comment:文件夹简介"`
	Name      string `gorm:"comment:文件夹名称"`
	Password  string `gorm:"comment:文件夹密码"`
	Size      int64  `gorm:"comment:文件夹大小"`
	Pid       string `gorm:"comment:文件夹父级ID"`
}
