package token

import (
	"crypto/rand"
	"fmt"
	"log"
)

func GenerateToken() string {
	key := make([]byte, 8)
	_, err := rand.Read(key)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%x", key)
}
