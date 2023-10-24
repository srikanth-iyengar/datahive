package router

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pipeline/{id}", PipelineHandler)
	r.HandleFunc("/worker/{id}", WorkerHandler)
	return r
}
