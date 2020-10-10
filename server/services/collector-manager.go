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
				//remove from
				err := n.ProcessorService.RemoveProcessor(&c)
				if err != nil {
					log.Println("Unable to remove ", err)
				}

				log.Println("Removed : ", c)
			}
		}
		log.Println("Waiting ...")
		time.Sleep(2 * time.Second)
	}

}

func (n *CollectorManager) HealthCheck(collector, collectorURL string) error {

	for i := 0; i < retries; i++ {
		resp, err := http.Get(collectorURL + "/health")

		if err != nil {
			log.Printf("%s is unhealthy : %v", collectorURL, err)
		} else if resp.StatusCode == 200 {
			return nil
		}

		time.Sleep(timeOut)
	}

	log.Printf("%s healthcheck failed after %d retries\n", collectorURL, retries)

	return fmt.Errorf("%s healthcheck failed after %d retries", retries)
}
