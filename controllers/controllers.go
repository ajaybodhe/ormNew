package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ajaybodhe/ormNew/models"
	"github.com/ajaybodhe/ormNew/queues"
	"github.com/ajaybodhe/ormNew/constants"
	"github.com/ajaybodhe/ormNew/conf"
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

func FakeHandler(msg []byte) error {
	return nil
}

func (this *MainController) Post() {
	o := orm.NewOrm()
    o.Using("default") //using default, you can use other database
	
	profile := new(models.Profile)
   	age,_ := strconv.Atoi(this.Input().Get("Age"))
	profile.Age = int16(age)
	
	user := new(models.User)
    user.Profile = profile
    user.Name = this.Input().Get("Name")
	
	fmt.Printf("body is %v\n", string(this.Ctx.Input.RequestBody))
	
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	
	que := queues.QueueCreation(constants.NSQ, FakeHandler, conf.OrmNewConfig.Queue.Topic)
	que.Publish(string(this.Ctx.Input.RequestBody))
	que.Stop()
}