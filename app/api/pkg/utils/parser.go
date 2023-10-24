package utils

import (
	"fmt"

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
	} `yaml:"kafka"`
}

func UpdateWorkers(config string, id string) []map[string]string {
	var conf DatahiveConfig
	err := yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	workers := []map[string]string{}
	for _, worker := range conf.KafkaWorkers {
		log.Info().Msg(fmt.Sprintf("Starting kafka consumer for inTopic: %s, hdfs: %t", worker.InTopic, worker.Hdfs))
		if worker.Hdfs {
			resp := StartKafkaWithHdfsPlugin(worker.InTopic, worker.HdfsFileName, id)
			w := map[string]string{
				"Id":         resp.GetWorkerId(),
				"Status":     resp.GetStatus().String(),
				"Type":       "Consumer",
				"PipelineId": id,
			}
			workers = append(workers, w)
		} else {
			resp := StartKafkaWithTransform(worker.InTopic, worker.Transform, worker.OutTopic, id)
			w := map[string]string{
				"Id":         resp.GetWorkerId(),
				"Status":     resp.GetStatus().String(),
				"Type":       "ConsumerWithHdfs",
				"PipelineId": id,
			}
			workers = append(workers, w)
		}
	}
	return workers
}
