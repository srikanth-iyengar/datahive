package utils

import (
    "os"

	"github.com/rs/zerolog/log"
    "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
)

type ServiceAddr string

var SparkAddr ServiceAddr = ServiceAddr(os.Getenv("SPARK_ADDR"))
var IngestionAddr ServiceAddr = ServiceAddr(os.Getenv("INGESTION_ADDR"))



func getNewGrpcClient(service_addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(service_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return conn, nil
}
