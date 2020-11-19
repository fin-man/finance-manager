package services

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type CollectorManager struct {
	ProcessorService *ProcessorService
}

var (
	retries = 2
	timeOut = 5 * time.Second
)

func NewCollectorManager(processorService *ProcessorService) *CollectorManager {

	return &CollectorManager{
		ProcessorService: processorService,
	}

}

func (n *CollectorManager) RunCollectorHealthChecks() {

	for {
		processors, err := n.ProcessorService.GetAllProcessors()

		if err != nil {
			log.Println("Unable to fetch registered collectors ", err)
		}

		for _, c := range processors {

			err = n.HealthCheck(c.ProcessorName, c.URL)
			if err != nil {
				log.Println(err)
				log.Println("Removing processor : ", c.ProcessorName)
				err = n.ProcessorService.RemoveProcessor(&c)
				if err != nil {
					log.Println(err)
				}
			}
		}
		time.Sleep(2 * time.Second)
	}

}

func (n *CollectorManager) HealthCheck(collector, collectorURL string) error {

	resp, err := http.Get(collectorURL + "/health")

	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("%s is unhealthy : %v", collectorURL, err)
	}

	return nil

}
