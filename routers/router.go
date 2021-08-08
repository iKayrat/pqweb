package routers

import (
	"pqweb/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/home", &controllers.MainController{})

	beego.Router("/", &controllers.MainController{}, "get:Get")

	beego.Router("/auth/login/:back", &controllers.MainController{}, "get,post:Login")
	beego.Router("/auth/register", &controllers.MainController{}, "get,post:SignUp")
	beego.Router("/user/logout", &controllers.MainController{}, "get:Logout")

	beego.Router("/notice", &controllers.MainController{}, "get:Notice")

	// beego.Router("/auth/login/:back", &controllers.Authorization{}, "get,post:Login")
	// beego.Router("/auth/register", &controllers.Authorization{}, "get,post:SignUp")
}
