package main

import (
	"beego-swagger-rest-api/database"
	_ "beego-swagger-rest-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//db
	database.InitDB()

	//----------------
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
