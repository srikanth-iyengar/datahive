package router

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pipeline/{id}", PipelineHandler)
	return r
}
