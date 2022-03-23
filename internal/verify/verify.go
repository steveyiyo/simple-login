package verify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Define Account Struct
type Account struct {
	Username string
	Password string
}

// Read User Data
func ReadUserData() []byte {
	filename := "data/account.json"
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

// Save the User Data
func SaveData(content Account) bool {
	filename := "data/account.json"
	saveData, _ := json.Marshal(content)
	err := ioutil.WriteFile(filename, saveData, os.ModeAppend)
	if err != nil {
		return false
	}
	return true
}

// Login Check
func Login(c *gin.Context) {
	var accountdata Account
	json.Unmarshal(ReadUserData(), &accountdata)
	input_username := c.PostForm("username")
	intput_password := c.PostForm("password")
	if input_username == accountdata.Username {
		if CheckPasswordHash(intput_password, accountdata.Password) {
			c.HTML(200, "index.tmpl", nil)
		} else {
			c.Data(400, "text/plain; charset=utf-8;", []byte("登入失敗！\n嘗試的使用者："+input_username))
		}
	} else {
		c.Data(400, "text/plain; charset=utf-8;", []byte("登入失敗！\n找不到指定的使用者："+input_username))
	}
}

// Register Check
func Register(c *gin.Context) {
	input_username := c.PostForm("username")
	HashPWD := HashPassword(c.PostForm("password"))

	user := Account{
		Username: input_username,
		Password: HashPWD,
	}

	if SaveData(user) {
		c.Data(200, "text/plain; charset=utf-8;", []byte("註冊成功！\n使用者："+input_username))
	} else {
		c.Data(400, "text/plain; charset=utf-8;", []byte("註冊失敗！\n請聯繫管理員"))
	}
}

// Hash password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if nil != err {
		fmt.Println(err)
	}
	return string(bytes)
}

// Check if password valid
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
