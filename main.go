package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Username string
	Password string
}

func webServer(c *gin.Context) {
	c.Redirect(302, "login")
}

func login_Page(c *gin.Context) {
	c.HTML(200, "login.tmpl", nil)
}

func register_Page(c *gin.Context) {
	c.HTML(200, "register.tmpl", nil)
}

func login_check(c *gin.Context) {
	var accountdata Account
	json.Unmarshal(ReadFile(), &accountdata)
	input_username := c.PostForm("username")
	intput_password := c.PostForm("password")
	if input_username == accountdata.Username {
		if CheckPasswordHash(intput_password, accountdata.Password) {
			// c.Data(200, "text/plain; charset=utf-8;", []byte("登入成功！\n使用者："+input_username))
			c.HTML(200, "index.tmpl", nil)
		} else {
			c.Data(400, "text/plain; charset=utf-8;", []byte("登入失敗！\n嘗試的使用者："+input_username))
		}
	} else {
		c.Data(400, "text/plain; charset=utf-8;", []byte("登入失敗！\n找不到指定的使用者："+input_username))
	}
}

func register_check(c *gin.Context) {
	input_username := c.PostForm("username")
	HashPWD := HashPassword(c.PostForm("password"))

	user := Account{
		Username: input_username,
		Password: HashPWD,
	}

	if SaveFile(user) {
		c.Data(200, "text/plain; charset=utf-8;", []byte("註冊成功！\n使用者："+input_username))
	} else {
		c.Data(400, "text/plain; charset=utf-8;", []byte("註冊失敗！\n請聯繫管理員"))
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if nil != err {
		fmt.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ReadFile() []byte {
	filename := "data/account.json"
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func SaveFile(content Account) bool {
	filename := "data/account.json"
	saveData, _ := json.Marshal(content)
	err := ioutil.WriteFile(filename, saveData, os.ModeAppend)
	if err != nil {
		return false
	}
	return true
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("static/*")

	router.GET("/", webServer)
	router.GET("/login", login_Page)
	router.GET("/register", register_Page)
	router.StaticFS("static/", http.Dir("./static_file"))

	router.POST("/login", login_check)
	router.POST("/register", register_check)

	router.NoRoute(pageNotAvailable)

	router.Run(":8082")
}
