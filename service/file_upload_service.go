package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"myfile/model"
	"myfile/util"
	"os"
	path2 "path"
	"strconv"
)

// FileUploadService 接收参数设置
type FileUploadService struct {
	file string `form:"file" json:"file"`
	hash string `form:"hash" json:"hash"`
	size string `form:"size" json:"size"`
}

func (service *FileUploadService) UploadFile(c *gin.Context) {
	fileobj := model.File{}
	//存储上传文件

	hash := c.Request.FormValue("hash")
	uid, _ := strconv.Atoi(c.Request.FormValue("uid"))
	fileExist, fileInfo := model.FastUploadFile(hash)
	if !fileExist {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Printf("failed to get data,err %s", err.Error())
			return
		}
		//fmt.Println(path2.Ext(header.Filename))

		path := "files/" + util.Md5Str(header.Filename) + path2.Ext(header.Filename)

		newFile, err := os.Create(path)
		if err != nil {
			fmt.Printf("failed to create data,err %s", err.Error())
			return
		}

		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("failed to copy data,err %s", err.Error())
			return
		}
		fileobj.Uid = int64(uid)
		fileobj.Path = path
		fileobj.FileName = header.Filename
		fileobj.SaveName = util.Md5Str(header.Filename) + path2.Ext(header.Filename)
		fileobj.Ext = path2.Ext(header.Filename)
		fileobj.Size = header.Size
		fileobj.Sha1 = hash

		if err = model.DB.Create(&fileobj).Error; err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	fileobj.Uid = int64(uid)
	fileobj.Path = fileInfo.Path
	fileobj.FileName = fileInfo.FileName
	fileobj.SaveName = fileInfo.SaveName
	fileobj.Ext = fileInfo.Ext
	fileobj.Size = fileInfo.Size
	fileobj.Sha1 = fileInfo.Sha1

	if err := model.DB.Create(&fileobj).Error; err != nil {
		fmt.Println(err)
		return
	}

}
