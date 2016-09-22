package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ajaybodhe/ormNew/models"
	"fmt"
)

type MainController struct {
        beego.Controller
}

func (this *MainController) Get() {
	o := orm.NewOrm()
    o.Using("default") //using default, you can use other database
	
	id,_:= strconv.Atoi(this.Ctx.Input.Param(":id"))
	user := models.User{Id: id}
	
	err := o.Read(&user)
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.Name)
	}
	
    this.Data["json"] = user
	this.Controller.ServeJSON()
}