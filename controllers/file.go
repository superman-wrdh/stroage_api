package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-common/app/service/main/member/model"
	"path"
	"strings"
)

// Operations about Users
type FileController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *FileController) Post() {
	uid := model.UUID4()
	uid = strings.Replace(uid, "-", "", -1)
	file, header, _ := u.GetFile("file")
	ContentType := header.Header.Get("Content-Type")
	FileSize := header.Size
	fmt.Println("file size : ", FileSize)
	fmt.Println("content type:", ContentType)
	ext := path.Ext(header.Filename)
	file.Close()
	fmt.Println("file ext:", ext)
	fmt.Print(header.Filename)
	SavePath := beego.AppConfig.String("StroageFilePath")
	fileFullPath := SavePath + uid + ext
	fmt.Println("save path:", fileFullPath)
	u.SaveToFile("file", fileFullPath)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}
