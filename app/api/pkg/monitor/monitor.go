package monitor

import (
	"time"

	"github.com/rs/zerolog/log"
	"datahive.io/api/internal/model"
	"datahive.io/api/pkg/utils"
)

func MonitorAll() {
	statusCheckTicker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-statusCheckTicker.C:
				statusCheckConsumer()
                statusCheckSparkJobs()
			}
		}
	}()
}

func statusCheckConsumer() {
	transformWorkers, _ := model.FindAllWorker(string(model.Consumer))
    hdfsWorkers, _ := model.FindAllWorker(string(model.ConsumerWithHdfs))
    workers := append(transformWorkers, hdfsWorkers...)
	for _, worker := range workers {
		if worker.Status == "NOT_FOUND" {
			continue
		}
		resp := utils.CheckConsumer(worker.Id)
		worker.Status = resp.GetStatus().String()
		worker.Update()
	}
}

func statusCheckSparkJobs() {
    sparkWorkers, _ := model.FindAllWorker(string(model.AnalysisJob))
    for _, worker := range sparkWorkers {
        if worker.Status == "NOT_FOUND" {
            continue
        }
        resp := utils.CheckSparkjob(worker.Id)
        worker.Status = resp.GetJobStatus().String()
        log.Info().Msgf("Worker Update status: %s, id, %s", worker.Id, worker.Status)
        worker.Update()
    }
}
