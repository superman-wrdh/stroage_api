package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"path"
	"stroage_api/models"
	"stroage_api/utils"
	"time"
)

// Operations about Users
type FileController struct {
	beego.Controller
}

type DownLoadController struct {
	beego.Controller
}

func ver(token string) bool {
	if token == "hcissuperman" {
		return true
	}
	return false
}

// @Title 单文件上传
// @Description create users
// @Param	body		body 	models.Resources	true		"body for user content"
// @Success 200 {int} models.Resources.Id
// @Failure 403 body is empty
// @router / [post]
func (u *FileController) Post() {
	description := u.GetString("description")
	token := u.GetString("token")
	IsAllow := ver(token)
	if !IsAllow {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Ctx.Output.Body([]byte("403 FORBIDDEN"))
		return
	}
	uid := utils.GetUUID()
	file, header, _ := u.GetFile("file")
	ContentType := header.Header.Get("Content-Type")
	FileSize := header.Size
	ext := path.Ext(header.Filename)
	file.Close()
	SavePath := beego.AppConfig.String("StorageFilePath")
	fileFullPath := SavePath + uid + ext
	u.SaveToFile("file", fileFullPath)
	r := models.Resources{
		Id:               uid,
		FileKey:          uid + ext,
		Type:             "",
		MimeType:         ContentType,
		ReferenceId:      "",
		Name:             uid + ext,
		OriginalFileName: header.Filename,
		Description:      description,
		Extension:        ext,
		StorageType:      "local",
		StorageParam:     "",
		Size:             FileSize,
		Meta:             "",
		CreatedTime:      time.Now(),
	}
	DataId, IsSuccess := r.Insert()
	fmt.Println(DataId, IsSuccess)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title 文件预览
// @Description get Resources by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Resources
// @Failure 404 :uid is empty
// @router :uid [get]
func (u *FileController) Get() {
	uid := u.GetString(":uid")
	r := models.Resources{}
	status := r.Get(uid)
	if !status {
		u.Ctx.ResponseWriter.WriteHeader(404)
		u.Ctx.Output.Body([]byte("404 NOT FOUND"))
		return
	}
	filePath := path.Join("file", r.FileKey)
	//fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(file.Name())
	file, _ = os.Open(filePath)
	u.Ctx.Output.Header("Content-Type", r.MimeType)
	//output.Header("Content-Disposition", "attachment; "+fn)
	//u.Ctx.Output.Header("Content-Disposition", "attachment; filename="+r.OriginalFileName)
	//u.Ctx.Output.Header("Content-Description", "File Transfer")
	fmt.Println(r.OriginalFileName)
	buffer, _ := ioutil.ReadAll(file)
	er := u.Ctx.Output.Body(buffer)
	if er != nil {

	}
	defer file.Close()
	//u.Ctx.Output.Download(filePath)

}

// @Title 文件下载
// @Description get Resources by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Resources
// @Failure 404 :uid is empty
// @router :uid [get]
func (u *DownLoadController) Get() {
	uid := u.GetString(":uid")
	r := models.Resources{}
	status := r.Get(uid)
	if !status {
		u.Ctx.ResponseWriter.WriteHeader(404)
		u.Ctx.Output.Body([]byte("404 NOT FOUND"))
		return
	}
	filePath := path.Join("file", r.FileKey)
	u.Ctx.Output.Download(filePath, r.OriginalFileName)
}

/**
    filePath := path.Join("file", "050ffec8689e463fa0741be0656188f7.jpg")
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())
	u.Ctx.Output.Header("Content-Type", "image/jpg")
	u.Ctx.Output.Download(filePath)
*/
