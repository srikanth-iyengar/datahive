package utils

import (
	"context"
	"os"
	"time"

	pb "datahive.io/api/pkg/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getNewGrpcClient() (*grpc.ClientConn, error) {
	ingestion_addr := os.Getenv("INGESTION_ADDR")
	conn, err := grpc.Dial(ingestion_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return conn, nil
}

func StartKafkaWithTransform(inTopic string, groovyScript string, outTopic string, pipelineId string) *pb.Response {
	conn, err := getNewGrpcClient()
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return &pb.Response{}
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.StartKafkaConsumerWithTransformation(ctx, &pb.ConsumerRequestWithTransformation{
		InTopic:      inTopic,
		GroovyScript: groovyScript,
		OutTopic:     outTopic,
		PipelineId:   pipelineId,
	})
	if err != nil {
		log.Err(err)
	}
	if resp.GetStatus() != pb.Response_RUNNING {
		log.Error().Msg("Error starting kafka consumer")
	}
	return resp
}

func StartKafkaWithHdfsPlugin(inTopic string, hdfsFileName string, pipelineId string) *pb.Response {
	conn, err := getNewGrpcClient()
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return &pb.Response{}
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.StartKafkaConsumerWithHDFSPlugin(ctx, &pb.ConsumerRequestWithHadoop{
		InTopic:      inTopic,
		HdfsFileName: hdfsFileName,
		PipelineId:   pipelineId,
	})
	if resp.GetStatus() != pb.Response_RUNNING {
		log.Error().Msg("Error starting kafka or connect to hdfs")
	}
	return resp
}

func CheckConsumer(worker_id string) *pb.Response {
	conn, err := getNewGrpcClient()
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.CheckConsumer(ctx, &pb.PingRequest{
		WorkerId: worker_id,
	})
	return resp
}

func StopConsumer(worker_id string) *pb.Response {
	conn, err := getNewGrpcClient()
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.StopKafkaConsumerWithTransformation(ctx, &pb.PingRequest{
		WorkerId: worker_id,
	})
	return resp
}
