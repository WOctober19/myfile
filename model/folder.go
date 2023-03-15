package model

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Uid       int64
	Introduce string
	Name      string
	Password  string
	Size      int64
	Pid       string
}
