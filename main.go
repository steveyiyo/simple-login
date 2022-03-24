package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/simple-login/internal/verify"
)

// Web Page
func webServer(c *gin.Context) {
	c.Redirect(302, "login")
}

func login_Page(c *gin.Context) {
	c.HTML(200, "login.tmpl", nil)
}

func register_Page(c *gin.Context) {
	c.HTML(200, "register.tmpl", nil)
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("static/*")

	router.GET("/", webServer)
	router.GET("/login", login_Page)
	router.GET("/register", register_Page)
	router.StaticFS("static/", http.Dir("./static_file"))

	router.POST("/login", verify.Login)
	router.POST("/register", verify.Register)

	router.NoRoute(pageNotAvailable)

	router.Run(":8082")
}
