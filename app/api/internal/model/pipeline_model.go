package model

import (
	"fmt"

	pb "datahive.io/api/pkg/grpc"
	"datahive.io/api/pkg/utils"
	"github.com/rs/zerolog/log"
)

type Pipeline struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Configuration string `json:"configuration"`
}

type TableName string

const (
	PipelineDb TableName = "datahive_pipeline"
)

func (p Pipeline) toSql() string {
	return fmt.Sprintf("INSERT INTO %s (id, name, configuration) VALUES ('%s', '%s', '%s');", PipelineDb, p.Id, p.Name, p.Configuration)
}

func FindPipeline(id string) Pipeline {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = '%s'", PipelineDb, id)
	pipeline := Pipeline{}
	err := db.QueryRow(query).Scan(&pipeline.Id, &pipeline.Name, &pipeline.Configuration)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return pipeline
}

func (p Pipeline) Save() (bool, string) {
	db := newConn()
	defer db.Close()
	query := p.toSql()
	err := db.QueryRow(query).Scan()
	if err != nil {
		return false, err.Error()
	}
	return true, "Saved pipeline successfully : {}"
}

func (p Pipeline) Delete() (bool, string) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", PipelineDb, p.Id)

	err := db.QueryRow(query).Scan()

	if err != nil {
		return false, fmt.Sprintf("Cannot delete the pipeline with id: %s", p.Id)
	}
	return true, fmt.Sprintf("Deleted the pipeline with id: %s", p.Id)
}

func (p Pipeline) Update(newPipeline Pipeline) (bool, string) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("UPDATE %s SET name='%s', configuration='%s' WHERE id='%s';", PipelineDb, newPipeline.Name, newPipeline.Configuration, p.Id)

	err := db.QueryRow(query).Scan()

	if err != nil {
		log.Error().Msg(err.Error())
	}
	go func() {
		workers := utils.UpdateWorkers(newPipeline.Configuration, newPipeline.Id)
		for _, worker := range workers {
			consType := Consumer
			if worker["Type"] != "Consumer" {
				consType = ConsumerWithHdfs
			}
			w := Worker{
				Id:         worker["Id"],
				Status:     pb.Response_WorkerStatus_value[worker["Status"]],
				Type:       consType,
				PipelineId: worker["PipelineId"],
			}
			w.Save()
		}
	}()
	return true, fmt.Sprintf("Updated pipeline with id: %s", p.Id)
}

func initPipelineSchema() bool {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id varchar, name varchar, configuration varchar, PRIMARY KEY(id));", PipelineDb)
	db.QueryRow(query).Scan()
	log.Info().Msg("Successfully created model for Pipeline")
	return true
}
