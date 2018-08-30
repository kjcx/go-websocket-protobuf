package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}
func (c *MainController) Get(){
	c.Data["json"] = "a"
	c.ServeJSON()
}
func main() {
	beego.Run()
	beego.Router("/", &MainController{})
}