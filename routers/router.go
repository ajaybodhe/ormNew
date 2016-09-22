package routers

import (
        "github.com/ajaybodhe/ormNew/controllers"
        "github.com/astaxie/beego"
)

func init() {
        beego.Router("/user/:id", &controllers.MainController{})
}