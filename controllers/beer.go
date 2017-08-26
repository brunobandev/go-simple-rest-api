package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/burnera/go-simple-rest-api/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	BeerController struct {
		session *mgo.Session
	}
)

func NewBeerController(s *mgo.Session) *BeerController {
	return &BeerController{s}
}

func (bc BeerController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	b := []models.Beer{}

	if err := bc.session.DB("cervejario").C("beers").Find(nil).All(&b); err != nil {
		w.WriteHeader(404)
		return
	}

	bj, _ := json.Marshal(b)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", bj)
}

func (bc BeerController) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	b := models.Beer{}

	if err := bc.session.DB("cervejario").C("beers").FindId(oid).One(&b); err != nil {
		w.WriteHeader(404)
		return
	}

	bj, _ := json.Marshal(b)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", bj)
}

func (bc BeerController) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b := models.Beer{}

	json.NewDecoder(r.Body).Decode(&b)

	b.Id = bson.NewObjectId()

	bc.session.DB("cervejario").C("beers").Insert(b)

	bj, _ := json.Marshal(b)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(201)

	fmt.Fprintf(w, "%s", bj)
}

func (bc BeerController) Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := bc.session.DB("cervejario").C("beers").RemoveId(oid); err != nil {
		w.WriteHeader(401)
		return
	}

	w.WriteHeader(200)
}
