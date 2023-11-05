package router

import (
	"encoding/json"
	"net/http"

	"datahive.io/api/internal/model"
)

func StackHandler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		stacks := model.FindAllStack()
		jsonPayload, _ := json.Marshal(stacks)
		w.Write(jsonPayload)
		w.WriteHeader(200)
	default:
		w.Write([]byte("This method is not supported"))
		w.WriteHeader(405)
	}
}
