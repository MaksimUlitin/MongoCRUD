package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maksimulitin/pkg/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (usCont *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	usMod := models.User{}
	if err := usCont.session.DB("db-users").C("users").FindId(oid).One(&usMod); err != nil {
		w.WriteHeader(404)
		return
	}

	usJson, err := json.Marshal(usMod)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-Type", "applecation/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s \n", usJson)

}
func (usCont *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usMod := models.User{}
	json.NewDecoder(r.Body).Decode(&usMod)
	usMod.Id = bson.NewObjectId()

	usCont.session.DB("db-users").C("users").Insert(usMod)

	usJson, err := json.Marshal(usMod)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-Type", "applecation/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s \n", &usJson)
}

func (usCont *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid := bson.ObjectIdHex(id)

	if err := usCont.session.DB("db-users").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}
	w.Header().Set("content-Type", "applecation/json")
	fmt.Fprint(w, "deleted user", oid, "\n")

}
