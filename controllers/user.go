package controllers

import (
	"beego-swagger-rest-api/models"
	"encoding/json"

	// "fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {

	// u.TplName = "index.tpl"
	// u.Data["inputFirstName"] = u.GetString("inputFirstName")
	// u.Data["inputLastName"] = u.GetString("inputLastName")
	// fmt.Println(u.Data["inputFirstName"])
	// fmt.Println(u.Data["inputLastName"])

	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"Message": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}
