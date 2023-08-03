package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
)

type userDomain struct {
	email    string
	name     string
	age      int8
	password string
}

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		age:      age,
		password: password,
	}
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
