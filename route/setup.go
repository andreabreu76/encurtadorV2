package route

import (
	"github.com/andreabreu76/encurtadorV2/handler/links"
	"github.com/andreabreu76/encurtadorV2/handler/utils"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", utils.Pong).Methods("GET")

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	linkPath := apiV1.PathPrefix("/links").Subrouter()

	linkPath.HandleFunc("/new", links.InsertNewLink).Methods("POST")

	return r
}
