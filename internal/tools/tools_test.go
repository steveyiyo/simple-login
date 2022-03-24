package tools_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Generates a random string of a given length
func RandomString(length int) string {
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJULMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return (output.String())
}

func TestMain(t *testing.T) {
	fmt.Println(RandomString(15))
}
