package jwt

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	token, _ := GenerateToken("steveyi")
	fmt.Println(token)
	success, message := ValidateToken(token)
	fmt.Println(success, message)
}
