package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
}

var authUserDB = map[string]authUser{}

type authUser struct {
	email        string
	passwordHash string
}

var DefaultUserService userService

type userService struct {
}

// VerifyUser autentica al usuario
func (userService) VerifyUser(user User) bool {
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password))
	return err == nil
}

// CreateUser se da alta un usuario si este no existe
func (userService) CreateUser(newUser User) error {
	_, ok := authUserDB[newUser.Email]
	if ok {
		fmt.Println("el usuario ya existe")
		return errors.New("el usuario ya existe")
	}
	passwordHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return err
	}
	newAuthUser := authUser{
		email:        newUser.Email,
		passwordHash: passwordHash,
	}
	fmt.Printf("Usuario creado: %v", newAuthUser)
	authUserDB[newAuthUser.email] = newAuthUser
	return nil
}

// getPasswordHash genera un hash para el password
func getPasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
