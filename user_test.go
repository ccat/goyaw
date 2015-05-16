package goyaw

import (
	//"html/template"
	"github.com/ccat/goyaw"
	"os"
	"testing"
)

func Test1_Create_Auth(t *testing.T) {
	_, err := os.Stat("test.db")
	if err == nil {
		err = os.Remove("test.db")
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	goyawInst := goyaw.NewGoyawInstance(&goyaw.UserDBconfig{Type: "sqlite3", Config: "test.db"})

	err = goyawInst.UserDB.CreateUser("alice", "password")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = goyawInst.UserDB.CreateUser("alice", "password")
	if err == nil {
		t.Errorf("Multiple same user created")
	}

	err = goyawInst.UserDB.Auth("alice", "password")
	if err != nil {
		t.Errorf(err.Error())
	}
	err = goyawInst.UserDB.Auth("alice", "password2")
	if err == nil {
		t.Errorf("User authed by wrong password")
	}

	err = goyawInst.UserDB.Auth("alice2", "password")
	if err == nil {
		t.Errorf("User authed by wrong username")
	}

	err = goyawInst.UserDB.Auth("alice", "or 1=1")
	if err == nil {
		t.Errorf("User authed by SQL injection")
	}
}
