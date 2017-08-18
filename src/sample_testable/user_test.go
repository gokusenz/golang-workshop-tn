package main

import (
	"fmt"
	"testing"
)

type MockUser struct{}

func (mu MockUser) CreateUser() User {
	fmt.Println("Call from mock user")
	return User{}
}

func TestCreateDataWithSuccess(t *testing.T) {
	register(MockUser{})
}
