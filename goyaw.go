package goyaw

import (
	//"fmt"
	//"github.com/codegangsta/negroni"
	//"html/template"
	//"net/http"
	//"os"
	Goji "github.com/zenazn/goji"
	//Graceful "github.com/zenazn/goji/graceful"
	//Web "github.com/zenazn/goji/web"
	//"golang.org/x/crypto/bcrypt"
	//"github.com/naoina/genmai"
)

//https://godoc.org/github.com/zenazn/goji
//https://godoc.org/golang.org/x/crypto/bcrypt

type GoyawInstance struct {
	//Mux      *http.ServeMux
	//Addr     string
	//CertFile string
	//KeyFile  string
	UserDB *UserMgmt
}

type UserDBconfig struct {
	Type   string
	Config string
}

func NewGoyawInstance(userDBconfig *UserDBconfig) *GoyawInstance {
	var yawIns *GoyawInstance = new(GoyawInstance)
	if userDBconfig != nil {
		yawIns.UserDB = NewUserDB(userDBconfig.Type, userDBconfig.Config)
	} else {
		yawIns.UserDB = nil
	}
	return yawIns
}

func (self *GoyawInstance) Serve() {
	Goji.Serve()
}

/*
https://github.com/zenazn/goji/issues/40


package main

import (
    "net/http"
    "fmt"

    "github.com/zenazn/goji/graceful"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
)

func main() {

    r := web.New()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/", IndexHandler)

    graceful.ListenAndServeTLS(":8000", "cert.pem", "key.pem", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", "world")
}
*/
