package model

import "github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
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

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string, age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
