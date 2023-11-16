package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maksimulitin/internal/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	usCont := controllers.NewUserController(getSession())
	r.GET("/user/:id", usCont.GetUser)
	r.POST("/user", usCont.CreateUser)
	r.DELETE("/user/:id", usCont.DeleteUser)
	http.ListenAndServe(":8088", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		log.Fatal(err)
	}
	return session
}
