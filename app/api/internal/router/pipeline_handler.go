package router

import (
	"datahive.io/api/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func PipelineHandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := io.ReadAll(r.Body)
	vars := mux.Vars(r)
	defer r.Body.Close()
	body, _ := unmarshallBody(bytes)

	switch method := r.Method; method {
	case "GET":
		status, respBody := getPipeline(vars["id"])
		w.Write(respBody)
		w.WriteHeader(status)
	case "DELETE":
		status, respBody := deletePipeline(vars["id"])
		w.Write(respBody)
		w.WriteHeader(status)
	case "POST":
		status, respBody := savePipeline(body)
		w.Write(respBody)
		w.WriteHeader(status)
	case "PUT":
		status, respBody := updatePipeline(body)
		w.Write(respBody)
		w.WriteHeader(status)
	default:
		w.Write([]byte(fmt.Sprintf("%s not supported", r.Method)))
	}
}

func updatePipeline(req model.Pipeline) (int, []byte) {
	pipeline := model.FindPipeline(req.Id)
	success, msg := pipeline.Update(req)
	if !success {
		log.Error().Msg(msg)
		return 500, []byte{}
	}
	resp, err := marshallBody(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return 500, []byte{}
	}
	return 200, resp
}

func getPipeline(id string) (int, []byte) {
	pipeline := model.FindPipeline(id)
	resp, err := marshallBody(pipeline)
	if err != nil {
		log.Error().Msg(err.Error())
		return 500, []byte{}
	}
	return 200, resp
}

func savePipeline(req model.Pipeline) (int, []byte) {
	success, msg := req.Save()
	if !success {
		log.Error().Msg(msg)
		return 500, []byte{}
	}
	resp, err := marshallBody(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return 500, []byte{}
	}
	return 200, resp
}

func deletePipeline(id string) (int, []byte) {
	pipeline := model.FindPipeline(id)
	success, msg := pipeline.Delete()
	if !success {
		log.Error().Msg(msg)
		return 500, []byte{}
	}
	resp, err := marshallBody(pipeline)
	if err != nil {
		log.Error().Msg(err.Error())
		return 500, []byte{}
	}
	return 200, resp
}

func unmarshallBody(bytes []byte) (model.Pipeline, error) {
	var pipeline model.Pipeline
	if err := json.Unmarshal(bytes, &pipeline); err != nil {
		log.Error().Msg(err.Error())
		return model.Pipeline{}, err
	}
	return pipeline, nil
}

func marshallBody(pipeline model.Pipeline) ([]byte, error) {
	body, err := json.Marshal(pipeline)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
