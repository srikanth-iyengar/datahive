package model

import (
	"fmt"

	"datahive.io/api/pkg/utils"
	"github.com/rs/zerolog/log"
)

type Pipeline struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Configuration string `json:"configuration"`
}

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
	db.QueryRow(query).Scan(&pipeline.Id, &pipeline.Name, &pipeline.Configuration)
	return pipeline
}

func FindAllPipeline() []Pipeline {
    db := newConn()
    defer db.Close()
    query := fmt.Sprintf("SELECT * FROM %s;", PipelineDb)
    resp := []Pipeline{}
    rows, _ := db.Query(query)
    for rows.Next() {
        var pipeline Pipeline
        rows.Scan(&pipeline.Id, &pipeline.Name, &pipeline.Configuration)
        resp = append(resp, pipeline)
    }
    return resp
}

func (p Pipeline) Save() (bool, string) {
	db := newConn()
	defer db.Close()
	query := p.toSql()
	db.Exec(query)
	return true, "Saved pipeline successfully : {}"
}

func (p Pipeline) Delete() (bool, string) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", PipelineDb, p.Id)

	db.Exec(query)

	return true, fmt.Sprintf("Deleted the pipeline with id: %s", p.Id)
}

func (p Pipeline) Update(newPipeline Pipeline) (bool, string) {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf("UPDATE %s SET name='%s', configuration='%s' WHERE id='%s';", PipelineDb, newPipeline.Name, newPipeline.Configuration, p.Id)

	db.Exec(query)

	go func() {
		workers := utils.UpdateWorkers(newPipeline.Configuration, newPipeline.Id)
		for _, worker := range workers {
			w := Worker{
				Id:         worker["Id"],
				Status:     worker["Status"],
				Type:       WorkerType(worker["Type"]),
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
	db.Exec(query)
	query = fmt.Sprintf("ALTER TABLE %s ALTER COLUMN id SET NOT NULL;", PipelineDb)
	db.Exec(query)
	log.Info().Msg("Init schema for pipeline ✅")
	return true
}
