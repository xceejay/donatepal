package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/routes"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/**/**/*.html")
	r.Static("/css", "views/css")
	r.Static("/images", "views/images")

	routes.InitRouter(r)

	r.Run(":8080")
}
