package utils

import (
	"context"
	"time"

	pb "datahive.io/api/pkg/grpc"
	"github.com/rs/zerolog/log"
)

func  StartSparkJob(jobName string,
                    driverMemory string,
                    executorMemory string,
                    resourceLocation string,
                    mainClass string) (*pb.JobInfo) {
    conn, err := getNewGrpcClient(string(SparkAddr))
    defer conn.Close()
    if err != nil {
        log.Error().Msg(err.Error())
        return &pb.JobInfo{
            JobStatus: pb.JobInfo_SCHEDULING_FAILURE,
        }
    }
    c := pb.NewAnalysisServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    defer cancel()

    resp, err := c.StartSparkJob(ctx, &pb.SparkJobRequest {
        JobName: jobName,
        DriverMemory: driverMemory,
        ExecutorMemory: executorMemory,
        ResourceLocation: resourceLocation,
        MainClass: mainClass,
    })
    if err != nil {
        log.Err(err)
    }
    return resp
}

func CheckSparkjob(workerId string) (*pb.JobInfo) {
    conn, err := getNewGrpcClient(string(SparkAddr))
    defer conn.Close()
    if err != nil {
        log.Error().Msg(err.Error())
        return &pb.JobInfo {
            JobStatus: pb.JobInfo_SCHEDULING_FAILURE,
        }
    }
    c := pb.NewAnalysisServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    resp, err := c.CheckJob(ctx, &pb.PingJob{
        WorkerId: workerId,
    })
    if err != nil {
        log.Error().Msg(err.Error())
    }
    return resp
}

func StopSparkJob(workerId string) (*pb.JobInfo) {
    conn, err := getNewGrpcClient(string(SparkAddr))
    defer conn.Close()
    if err != nil {
        log.Error().Msg(err.Error())
        return &pb.JobInfo {
            JobStatus: pb.JobInfo_SCHEDULING_FAILURE,
        }
    }
    c := pb.NewAnalysisServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    resp, err := c.StopSparkJob(ctx, &pb.PingJob{
        WorkerId: workerId,
    })
    if err != nil {
        log.Error().Msg(err.Error())
    }
    return resp
}
