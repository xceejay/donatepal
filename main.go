package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xceejay/boilerplate/routes"
)

func main() {

	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.LoadHTMLGlob("views/**/**/*.html")
	r.Static("/css", "views/css")
	r.Static("/images", "views/images")
	r.Static("/ext", "views/ext")
	r.Static("/js", "views/js")

	routes.InitRouter(r)

	r.Run(":8080")
}
