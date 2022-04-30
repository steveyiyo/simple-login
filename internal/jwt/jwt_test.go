package jwt

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	token, _ := GenerateToken("steveyi")
	fmt.Printf("Token: %s\n\n", token)

	_, message := ValidateToken(token)
	fmt.Printf("Status: %s\n", message)
}
