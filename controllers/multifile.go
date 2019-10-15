package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"stroage_api/models"
	"stroage_api/utils"
	"time"
)

type MultiFileController struct {
	beego.Controller
}

//保存文件
func SaveFile(FileFullPath string, file multipart.File) bool {
	defer file.Close()
	dst, err := os.Create(FileFullPath)
	if err != nil {
		return false
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return false
	}
	return true
}

// @Title 多文件上传
// @Description create users
// @Param	body		body 	models.Resources	true		"body for user content"
// @Success 200 {int} models.Resources.Id
// @Failure 403 body is empty
// @router / [post]
func (u *MultiFileController) Post() {
	description := u.GetString("description")
	token := u.GetString("token")
	IsAllow := ver(token)
	IsAllow = true
	if !IsAllow {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Ctx.Output.Body([]byte("403 FORBIDDEN"))
		return
	}

	files, err := u.GetFiles("file")
	if err != nil {
		http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusNoContent)
		return
	}

	FileIdList := make([]string, 5, 5)
	for index, file := range files {
		uid := utils.GetUUID()
		fmt.Println(index, file.Filename)
		ContentType := file.Header.Get("Content-Type")
		FileSize := file.Size
		ext := path.Ext(file.Filename)

		SavePath := beego.AppConfig.String("StorageFilePath")
		fileFullPath := SavePath + uid + ext
		IoFile, err := file.Open()
		if err != nil {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		IsSuccess := SaveFile(fileFullPath, IoFile)
		if !IsSuccess {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		FileIdList = append(FileIdList, uid)
		r := models.Resources{
			Id:               uid,
			FileKey:          uid + ext,
			Type:             "",
			MimeType:         ContentType,
			ReferenceId:      "",
			Name:             uid + ext,
			OriginalFileName: file.Filename,
			Description:      description,
			Extension:        ext,
			StorageType:      "local",
			StorageParam:     "",
			Size:             FileSize,
			Meta:             "",
			CreatedTime:      time.Now(),
		}
		DataId, IsSuccess := r.Insert()
		if !IsSuccess {
			fmt.Println("保存文信息到数据库件失败")
			fmt.Println(file.Filename)
			fmt.Println("保存文件失败")
		} else {
			fmt.Println("保存文件信息到数据库失败")
			fmt.Println(DataId, file.Filename)
			fmt.Println("保存文件信息到数据库失败")
		}
	}
	u.Data["json"] = map[string][]string{"files": FileIdList}
	u.ServeJSON()
	return
}
