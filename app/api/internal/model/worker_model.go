package model

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type WorkerType string

type tableName string

const (
	Consumer         WorkerType = "KafkaConsumer"
	ConsumerWithHdfs WorkerType = "KafkaConsumerWithHdfs"
	AnalysisJob      WorkerType = "SparkJob"

	schema tableName = "datahive_workers"
)

type Worker struct {
	Id         string     `json:"id"`
	Status     int32      `json:"status"`
	Type       WorkerType `json:"type"`
	PipelineId string     `json:"pipelineId"`
}

func (w Worker) Save() error {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO %s (id, status, type, pipeline_id) values('%s', %d, '%s', '%s');", schema, w.Id, w.Status, w.Type, w.PipelineId)
	err := db.QueryRow(query).Scan()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return nil
}

func (w Worker) Update() error {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("UPDATE %s SET status=%d, type='%s'  WHERE id='%s' and pipeline_id='%s';", schema, w.Status, w.Type, w.Id, w.PipelineId)
	err := db.QueryRow(query).Scan()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return nil
}

func FindWorker(id string) (Worker, error) {
	db := newConn()
	defer db.Close()
	var w Worker
	query := fmt.Sprintf("SELECT * FROM %s WHERE id='%s';", schema, id)
	err := db.QueryRow(query).Scan(w.Id, w.Status, w.Type, w.PipelineId)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return w, nil
}

func FindAllWorker() ([]Worker, error) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM %s;", schema)
	rows, err := db.Query(query)
	if err != nil {
		log.Warn().Msg(err.Error())
	}
	workers := []Worker{}
	defer rows.Close()
	for rows.Next() {
		var workerId, workerType, pipelineId string
		var status int32
		rows.Scan(&workerId, &status, &workerType, &pipelineId)
		w := Worker{
			Id:         workerId,
			Status:     status,
			Type:       WorkerType(workerType),
			PipelineId: pipelineId,
		}
		workers = append(workers, w)
	}
	return workers, nil
}

func FindWorkersByPipelineId(pipelineId string) ([]Worker, error) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM %s WHERE pipeline_id='%s';", schema, pipelineId)
	rows, err := db.Query(query)
	if err != nil {
		log.Warn().Msg(err.Error())
	}
	workers := []Worker{}
	defer rows.Close()
	for rows.Next() {
		var workerId, workerType, pipelineId string
		var status int32
		rows.Scan(&workerId, &status, &workerType, &pipelineId)
		w := Worker{
			Id:         workerId,
			Status:     status,
			Type:       WorkerType(workerType),
			PipelineId: pipelineId,
		}
		workers = append(workers, w)
	}
	return workers, nil
}

func initWorkerSchema() bool {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id varchar, status int, type varchar, pipeline_id varchar, PRIMARY KEY (id),CONSTRAINT fk_pipeline FOREIGN KEY(pipeline_id) REFERENCES %s(id));", schema, PipelineDb)
	err := db.QueryRow(query).Scan()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Info().Msg("Init schema for worker successfull")
	return true
}
