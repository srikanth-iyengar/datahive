package utils

import (
	"datahive.io/api/pkg/grpc"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type DatahiveConfig struct {
	IsStream     string `yaml:"type"`
	KafkaWorkers []struct {
		Hdfs         bool   `yaml:"hdfs,omitempty"`
		InTopic      string `yaml:"inTopic,omitempty"`
		OutTopic     string `yaml:"outTopic,omitempty"`
		Transform    string `yaml:"transform,omitempty"`
		HdfsFileName string `yaml:"hdfsFileName,omitempty"`
	} `yaml:"kafka,omitempty"`
	SparkWorkers []struct {
        JobName        string `yaml:"job-name"`
        AppResource    string `yaml:"app-resource"`
        DriverMemory   string `yaml:"driver-memory"`
        ExecutorMemory string `yaml:"executor-memory"`
        ResLocation    string `yaml:"resouce-location"`
        MainClass      string `yaml:"main-class"`
    } `yaml:"spark,omitempty"`
}

func UpdateWorkers(config string, id string) []map[string]string {
	var conf DatahiveConfig
	err := yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	workers := []map[string]string{}
	for _, worker := range conf.KafkaWorkers {
		if worker.Hdfs {
			resp := StartKafkaWithHdfsPlugin(worker.InTopic, worker.HdfsFileName, id)
            if resp.GetStatus() != grpc.Response_RUNNING {
                log.Warn().Msgf("Failed to start kafka worker: %s, status: %s, msg: %s", resp.GetWorkerId(), resp.GetStatus().String(), resp.GetMessage())
                continue
            }
			w := map[string]string{
				"Id":         resp.GetWorkerId(),
				"Status":     resp.GetStatus().String(),
				"Type":       "KafkaConsumerWithHdfs",
				"PipelineId": id,
			}
			workers = append(workers, w)
		} else {
			resp := StartKafkaWithTransform(worker.InTopic, worker.Transform, worker.OutTopic, id)
            if resp.GetStatus() != grpc.Response_RUNNING {
                log.Warn().Msgf("Failed to start kafka worker: %s, status: %s, msg: %s", resp.GetWorkerId(), resp.GetStatus().String(), resp.GetMessage())
                continue
            }
			w := map[string]string{
				"Id":         resp.GetWorkerId(),
				"Status":     resp.GetStatus().String(),
				"Type":       "KafkaConsumer",
				"PipelineId": id,
			}
			workers = append(workers, w)
		}
	}
    
    for _, worker := range conf.SparkWorkers {
        resp := StartSparkJob(worker.JobName, worker.DriverMemory, worker.ExecutorMemory, worker.AppResource, worker.MainClass)
        if resp.GetJobStatus() != grpc.JobInfo_RUNNING {
            log.Warn().Msgf("Failed to start kafka worker: %s, status: %s, msg: %s", resp.GetWorkerId(), resp.GetJobStatus().String(), resp.GetMessage())
            continue
        }
        w := map[string]string {
            "Id": resp.GetWorkerId(),
            "Status": resp.GetJobStatus().String(),
            "Type": string("SparkJob"),
            "PipelineId": id,
        }
        workers = append(workers, w)
    }
	return workers
}
