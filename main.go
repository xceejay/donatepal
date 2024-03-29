package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"

	"github.com/xceejay/donatepal/routes"
)

func main() {

	r := gin.New()

	r.Use(sessions.Sessions("session", sessions.NewCookieStore([]byte("secret"))))

	r.Use(favicon.New("./favicon.ico"))
	r.LoadHTMLGlob("views/**/**/*.html")
	r.Static("/css", "views/css")
	r.Static("/images", "views/images")
	r.Static("/ext", "views/ext")
	r.Static("/js", "views/js")

	routes.InitRouter(r)

	r.Run(":8080")
}
