package main

import (
	"net/http"

	"github.com/burnera/go-simple-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := httprouter.New()

	bc := controllers.NewBeerController(getSession())

	r.GET("/beer", bc.Index)
	r.POST("/beer", bc.Store)
	r.GET("/beer/:id", bc.Show)
	r.DELETE("/beer/:id", bc.Destroy)

	http.ListenAndServe(":12345", r)
}
