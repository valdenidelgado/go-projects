package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	email    string
	name     string
	age      int8
	password string
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
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

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
