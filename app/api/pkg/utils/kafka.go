package utils

import (
	"context"
	"time"

	pb "datahive.io/api/pkg/grpc"
	"github.com/rs/zerolog/log"
)

func StartKafkaWithTransform(inTopic string, groovyScript string, outTopic string, pipelineId string) *pb.Response {
	conn, err := getNewGrpcClient(string(IngestionAddr))
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return &pb.Response{}
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
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
	conn, err := getNewGrpcClient(string(IngestionAddr))
	defer conn.Close()
	if err != nil {
		log.Error().Msg(err.Error())
		return &pb.Response{}
	}
	c := pb.NewIngestionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	conn, err := getNewGrpcClient(string(IngestionAddr))
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
	conn, err := getNewGrpcClient(string(IngestionAddr))
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
