package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// *****************************
// Public Functions
// *****************************
// These call the private functions for the purposes of abstraction
// No encapsulation via a struct/object b/c such an object would
// probably hold information similar to a Session object & therefore
// doesn't add anything practical to implementation of authentication

// AuthHashPassword returns Hashed version of input password
func AuthHashPassword(password string) (string, error) {
	return hash(password)
}

// AuthVerifyPassword verifies whether the input password matches the input hashed password string
func AuthVerifyPassword(password string, hashedPassword string) error {
	return verify(password, hashedPassword)
}

// *****************************
//  Private Functions
// *****************************

func hash(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), err
}

func verify(password, hashedPassword string) error {

	// return the result of CompareHashAndPassword
	// deal with the result wherever VerifyPassword is being called
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
