package model

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type Stack struct {
	Id              int32  `json:"id"`
	Name            string `json:"name"`
	Addr            string `json:"address"`
	EarliestSuccess int64  `json:"earliestSuccess"`
	IsUp            bool   `json:"isUp"`
}

const (
	StackDb      TableName = "datahive_stack"
	StackInsert  string    = "INSERT INTO %s (id, addr, earliest_success, name) VALUES (%d, '%s', %d, '%s');"
	StackUpdate  string    = "UPDATE %s SET addr='%s', earliest_success=%d, is_up=%t WHERE id=%d;"
	StackRead    string    = "SELECT * FROM %s WHERE id=%d;"
	StackReadAll string    = "SELECT * FROM %s;"
)

func (st Stack) Save() {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf(StackInsert, StackDb, st.Id, st.Addr, st.EarliestSuccess, st.Name)
	db.QueryRow(query).Scan()
}

func (st Stack) Update() {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf(StackUpdate, StackDb, st.Addr, st.EarliestSuccess, st.IsUp, st.Id)
	db.QueryRow(query).Scan()
}

func FindStackById(id int32) Stack {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf(StackRead, StackDb, id)
	var st Stack
	db.QueryRow(query).Scan(st.Id, st.Addr, st.EarliestSuccess, st.IsUp)
	return st
}

func FindAllStack() []Stack {
	db := newConn()
	defer db.Close()
	query := fmt.Sprintf(StackReadAll, StackDb)
	rows, _ := db.Query(query)
	result := []Stack{}
	defer rows.Close()
	for rows.Next() {
		var st Stack
		rows.Scan(&st.Id, &st.Addr, &st.EarliestSuccess, &st.IsUp, &st.Name)
		result = append(result, st)
	}
	return result
}

func InitStackSchema() {
	db := newConn()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id int, addr varchar, earliest_success int, is_up boolean, name varchar, PRIMARY KEY (id) );", StackDb)
	err := db.QueryRow(query).Scan()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	db.Close()
	stacks := []map[string]string{
		{
			"id":   "1",
			"addr": "SPARK_ADDR",
			"name": "Apache Spark",
		},
		{
			"id":   "2",
			"addr": "INGESTION_ADDR",
			"name": "Datahive Ingestor",
		},
		{
			"id":   "3",
			"addr": "NAMENODE_ADDR",
			"name": "Apache Hadoop",
		},
		{
			"id":   "4",
			"addr": "MINIO_ADDR",
			"name": "MinIo",
		},
		{
			"id":   "5",
			"addr": "ELASTIC_ADDR",
			"name": "Elastic",
		},
		{
			"id":   "6",
			"addr": "KIBANA_ADDR",
			"name": "Kibana",
		},
		{
			"id":   "7",
			"addr": "API_ADDR",
			"name": "Datahive Api",
		},
		{
			"id":   "8",
			"addr": "KAFKA_ADDR",
			"name": "Apache Kafka",
		},
	}
	for _, stack := range stacks {
		stack_addr := os.Getenv(stack["addr"])
		stack_name := stack["name"]
		stack_id, err := strconv.Atoi(stack["id"])
		if err != nil {
			continue
		}
		st := Stack{
			Id:              int32(stack_id),
			Name:            stack_name,
			Addr:            stack_addr,
			EarliestSuccess: time.Now().Unix(),
			IsUp:            false,
		}
		st.Save()
	}
	log.Info().Msg("Init schema for stack successful ✅")
}
