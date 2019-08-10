package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type WordController struct {
	beego.Controller
}

// @Title Get
// @Description get word detail by word
// @Param	word	path 	string	true		"The word"
// @Success 200
// @Failure 400 :word is empty
// @router /:word [get]
func (u *WordController) Get() {
	word := u.GetString(":word")
	u.Data["json"] = map[string]string{"word": word}
	u.ServeJSON()
	u.ServeJSON()
}
