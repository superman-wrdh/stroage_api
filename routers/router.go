// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"stroage_api/controllers"
	"stroage_api/filters"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/file/",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
		beego.NSNamespace("/multiFile/",
			beego.NSInclude(
				&controllers.MultiFileController{},
			),
		),
		beego.NSNamespace("/download/",
			beego.NSInclude(
				&controllers.DownLoadController{},
			),
		),
		beego.NSNamespace("/word/",
			beego.NSInclude(
				&controllers.WordController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//注册所有的过滤器
	filters.RegisterAllFilter()
}
