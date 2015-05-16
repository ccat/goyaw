package main

import (
	"fmt"
	"github.com/ccat/goyaw"
)

func main() {
	goyawInst := goyaw.NewGoyawInstance(&goyaw.UserDBconfig{Type: "sqlite3", Config: "test.db"})
	err := goyawInst.UserDB.CreateUser("alice", "password")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = goyawInst.UserDB.Auth("alice", "password")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = goyawInst.UserDB.Auth("alice", "password2")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = goyawInst.UserDB.Auth("alice2", "password")
	if err != nil {
		fmt.Println(err.Error())
	}
}
