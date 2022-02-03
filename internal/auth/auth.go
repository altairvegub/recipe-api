package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(pwd string) (string, error){
	// generate hash and salt from a password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPass), nil
}

func VerifyPassword (hashedPwd string, plainTextPwd []byte) bool {
	// convert string to byte slice
	hashedPwdByte := []byte(hashedPwd)

	// compare hash and password
	err := bcrypt.CompareHashAndPassword(hashedPwdByte, plainTextPwd)
	if err == nil {
		fmt.Println("The passwords match")
		return true
	} else {
		fmt.Println("The passwords do not match")
		return false
	}
}