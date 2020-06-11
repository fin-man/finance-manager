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
	retries = 2
	timeOut = 5 * time.Second
)

func NewCollectorManager(collectorService *CollectorService) *CollectorManager {

	return &CollectorManager{
		CollectorService: collectorService,
	}

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
				//remove from
				_, err := n.CollectorService.RemoveCollector(c)
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
