package custom

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(in string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing")
	}
	return string(hashedPwd)
}
