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
	Status     string     `json:"status"`
	Type       WorkerType `json:"type"`
	PipelineId string     `json:"pipelineId"`
}

func (w Worker) Save() error {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO %s (id, status, type, pipeline_id) values('%s', '%s', '%s', '%s');", schema, w.Id, w.Status, w.Type, w.PipelineId)
	db.QueryRow(query).Scan()
	return nil
}

func (w Worker) Update() error {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("UPDATE %s SET status='%s', type='%s'  WHERE id='%s' and pipeline_id='%s';", schema, w.Status, w.Type, w.Id, w.PipelineId)
    db.QueryRow(query).Scan()
	return nil
}

func FindWorker(id string) (Worker, error) {
	db := newConn()
	defer db.Close()
	var w Worker
	query := fmt.Sprintf("SELECT * FROM %s WHERE id='%s';", schema, id)
	db.QueryRow(query).Scan(w.Id, w.Status, w.Type, w.PipelineId)
	return w, nil
}

func FindAllWorker(workerType string) ([]Worker, error) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM %s WHERE type='%s';", schema, workerType)
	rows, _ := db.Query(query)
	workers := []Worker{}
	defer rows.Close()
	for rows.Next() {
		var workerId, workerType, pipelineId, status string
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
		var status string
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
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id varchar, status varchar, type varchar, pipeline_id varchar, PRIMARY KEY (id),CONSTRAINT fk_pipeline FOREIGN KEY(pipeline_id) REFERENCES %s(id));", schema, PipelineDb)
	db.QueryRow(query).Scan()
	query = fmt.Sprintf("ALTER TABLE %s ALTER COLUMN id SET NOT NULL;", PipelineDb)
	db.QueryRow(query).Scan()
	log.Info().Msg("Init schema for worker ✅")
	return true
}
