package model

import (
	"database/sql"

	"datahive.io/api/internal/config"
)

type TableName string

func newConn() *sql.DB {
	db := config.ConnectDb()
	return db
}

func InitSchema() {
	initPipelineSchema()
	initWorkerSchema()
	InitStackSchema()
}
