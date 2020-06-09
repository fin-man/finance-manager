package services

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type CollectorManager struct {
	CollectorService *CollectorService
}

var (
	retries = 3
	timeOut = 5 * time.Second
)

func NewCollectorManager() *CollectorManager {

	return &CollectorManager{}

}

func (n *CollectorManager) RunCollectorHealthChecks() {

	for {
		collectors, err := n.CollectorService.GetAllCollectors()

		if err != nil {
			log.Println("Unable to fetch registered collectors ", err)
		}

		for _, c := range collectors {
			collectorURL, err := n.CollectorService.GetNewCollector(c)
			if err != nil {
				log.Printf("Unable to fetch : %s collector \n", c)
				continue
			}

			err = n.HealthCheck(c, collectorURL)
			if err != nil {

			}
		}
	}

}

func (n *CollectorManager) HealthCheck(collector, collectorURL string) error {

	for i := 0; i < retries; i++ {
		resp, err := http.Get(collectorURL + "/healthcheck")

		if err != nil {
			log.Printf("%s is unhealthy : %v", err)
		} else if resp.StatusCode == 200 {
			break
		}

		time.Sleep(timeOut)
	}

	log.Printf("%s healthcheck failed after %d retries\n", retries)

	return fmt.Errorf("%s healthcheck failed after %d retries", retries)
}
