package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/simple-login/internal/jwt"
)

func Init() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("static/*")
	router.StaticFS("static/", http.Dir("./static_file"))

	router.GET("/", index)
	router.GET("/login", login_Page)
	router.GET("/register", register_Page)
	router.GET("/jwt_test", jwtTest)

	router.POST("/login", authLogin)
	router.POST("/register", authRegister)

	router.NoRoute(pageNotAvailable)

	router.Run(":8082")
}

// Web Page
func index(c *gin.Context) {
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

func jwtTest(c *gin.Context) {
	var token string
	token = ""

	cookies := c.Request.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "jwt" {
			token = cookie.Value
		}
	}

	if token == "" {
		c.String(400, "Unauthorized")
	} else {
		jwt_check, _ := jwt.ValidateToken(token)
		if jwt_check {
			c.String(200, "Authorized")
		} else {
			c.String(400, "Unauthorized")
		}
	}
}
