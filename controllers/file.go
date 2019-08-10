package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-common/app/service/main/member/model"
	"io/ioutil"
	"os"
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
	SavePath := beego.AppConfig.String("StorageFilePath")
	fileFullPath := SavePath + uid + ext
	fmt.Println("save path:", fileFullPath)
	u.SaveToFile("file", fileFullPath)
	u.Data["json"] = map[string]string{"uid": uid}
	u.Ctx.Output.Body([]byte("hello"))
	//u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [GET]
func (u *FileController) Get() {
	//appPath := beego.AppPath
	filePath := path.Join("file", "050ffec8689e463fa0741be0656188f7.jpg")
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())
	file, _ = os.Open(filePath)
	u.Ctx.Output.Header("Content-Type", "image/jpg")
	buffer, _ := ioutil.ReadAll(file)
	u.Ctx.Output.Body(buffer)
	//u.Ctx.Output.Download(filePath)

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
