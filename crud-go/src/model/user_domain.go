package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
)

type UserDomain struct {
	Email    string
	Name     string
	Age      int8
	Password string
}

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Name:     name,
		Age:      age,
		Password: password,
	}
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
