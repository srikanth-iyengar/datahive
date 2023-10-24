package monitor

import (
	"time"

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
			}
		}
	}()
}

func statusCheckConsumer() {
	workers, _ := model.FindAllWorker()
	for _, worker := range workers {
		if worker.Status == 3 {
			continue
		}
		resp := utils.CheckConsumer(worker.Id)
		worker.Status = int32(resp.GetStatus())
		worker.Update()
	}
}
