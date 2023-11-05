package monitor

import (
	"net"
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
				statusCheckSparkJobs()
				stackMonitor()
			}
		}
	}()
}

func stackMonitor() {
	stacks := model.FindAllStack()
	for _, stack := range stacks {
		conn, err := net.Dial("tcp", stack.Addr)
		if err != nil {
			stack.IsUp = false
			stack.Update()
		} else {
			defer conn.Close()
			if !stack.IsUp {
				stack.EarliestSuccess = time.Now().Unix()
			}
			stack.IsUp = true
			stack.Update()
		}
	}
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
		worker.Update()
	}
}
