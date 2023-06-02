// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements add new user and password verification.

This package uses package bcrypt, witch implements Provos
and Mazières's bcrypt adaptive hashing algorithm
*/

package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hash for the password...
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hash), err
}

// CheckPasswordHash checks password by hash...
func CheckPasswordHash(password, hash string) bool {
	fmt.Println(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash)) //move err
	return err == nil
}
