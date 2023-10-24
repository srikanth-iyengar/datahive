package model

import (
	"database/sql"
	"datahive.io/api/internal/config"
)

func newConn() *sql.DB {
	db := config.ConnectDb()
	return db
}

func InitSchema() {
	initPipelineSchema()
	initWorkerSchema()
}
