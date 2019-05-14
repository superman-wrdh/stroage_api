package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
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
	uid := ""
	file, header, _ := u.GetFile("file")
	ext := path.Ext(header.Filename)
	file.Close()
	fmt.Print(ext)
	fmt.Print(header.Filename)
	SavePath := beego.AppConfig.String("StroageFilePath")
	fileFullPath := SavePath + "XX" + ext
	fmt.Print("save path ", fileFullPath)
	u.SaveToFile("file", fileFullPath)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}
