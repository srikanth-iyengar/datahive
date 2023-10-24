package router

import (
	"encoding/json"
	"net/http"

	"datahive.io/api/internal/model"
	"github.com/gorilla/mux"
)

func WorkerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	switch method := r.Method; method {
	case "GET":
		workers, _ := model.FindWorkersByPipelineId(id)
		jsonPayload, _ := json.Marshal(workers)
		w.Write(jsonPayload)
		w.WriteHeader(200)
	default:
		w.Write([]byte("This method is not supported"))
		w.WriteHeader(405)
	}
}
