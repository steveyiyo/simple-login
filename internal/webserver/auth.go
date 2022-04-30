package webserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/simple-login/internal/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Define Account Struct
type Account struct {
	Username string
	Password string
}

type Result struct {
	Success  bool
	Message  string
	Username string
	Token    string
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
func authLogin(c *gin.Context) {

	// Define account_data
	var account_data Account
	json.Unmarshal(ReadUserData(), &account_data)

	// Get Login username and password
	input_username := c.PostForm("username")
	intput_password := c.PostForm("password")

	// Global variable
	var return_result Result
	var token string

	// Check if username and password valid
	if input_username == account_data.Username {
		if CheckPasswordHash(intput_password, account_data.Password) {
			token, _ = jwt.GenerateToken(input_username)

			// return success message
			return_result = Result{true, "登入成功！", input_username, token}

			c.SetCookie("jwt", token, 86400, "/", "", false, true)
			c.JSON(200, return_result)
		} else {
			token = ""

			// return success message
			return_result = Result{true, "登入失敗！", input_username, token}
			c.JSON(403, return_result)
		}
	} else {
		token = ""

		// return success message
		log.Println("User not found.")
		return_result = Result{true, "登入失敗！", input_username, token}
		c.JSON(403, return_result)
	}
}

// Register Check
func authRegister(c *gin.Context) {

	// Get Register username and hash the password
	username := c.PostForm("username")
	HashPWD := HashPassword(c.PostForm("password"))

	// Define account_data
	account_data := Account{
		Username: username,
		Password: HashPWD,
	}

	// Global variable
	var return_result Result

	// Save the User Data, and return the result
	if SaveData(account_data) {
		return_result = Result{true, "註冊成功！", username, ""}
		c.JSON(200, return_result)
	} else {
		return_result = Result{true, "註冊失敗！請聯繫管理員", "", ""}
		c.JSON(400, return_result)
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
