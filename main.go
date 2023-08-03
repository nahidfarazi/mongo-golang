package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nahidfarazi/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)
func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSessions())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSessions() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}