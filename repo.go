package framework

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type appContext struct {
	db *mgo.Database
}

func (c *appContext) teaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := TeaRepo{c.db.C("teas")}
	tea, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(tea)
}

func (c *appContext) AllTeasHandler(w http.ResponseWriter, r *http.Request) {
	repo := TeaRepo{c.db.C("teas")}
	teas, err := repo.findAll()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(teas)
}

func (c *appContext) createTeaHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*TeaResource)
	repo := TeaRepo{c.db.C("teas")}
	err := repo.Create(&body.Data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(body)
}

// TeaResource contains a reference to a single resource of Tea
type TeaResource struct {
	Data Tea `json:"data"`
}

// TeaResources contains a slice of tea for transferrance to the JSON-API standard
type TeaResources struct {
	Data []Tea `json:"data"`
}

// TeaRepo contains a reference to the mongoDB Tea Collection
type TeaRepo struct {
	coll *mgo.Collection
}

// Tea is the representation of the Tea data
type Tea struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
}

// Find searches the TeaRepo for a Tea with the id passed to the method
func (r *TeaRepo) Find(id string) (TeaResource, error) {
	result := TeaResource{}
	err := r.coll.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TeaRepo) findAll() (TeaResources, error) {
	result := TeaResources{}
	err := r.coll.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}
	return result, err
}

// Create adds an additional Tea to the TeaRepo with the specified data in the parameter
func (r *TeaRepo) Create(tea *Tea) error {
	id := bson.NewObjectId()
	fmt.Println(id)
	_, err := r.coll.UpsertId(id, tea)
	if err != nil {
		return err
	}
	tea.ID = id
	return err
}
