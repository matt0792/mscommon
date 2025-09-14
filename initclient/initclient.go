package initclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/matt0792/mscommon/commonmodels"
)

func init() {
	providerLoc := os.Getenv("PROVIDER_LOCATION")
	endpoint := providerLoc + "/services"

	config := commonmodels.MicroserviceConfig{
		ServiceLocation: os.Getenv("SERVICE_LOCATION"),
		ServiceName:     os.Getenv("SERVICE_NAME"),
		ServiceID:       os.Getenv("SERVICE_ID"),
	}

	payload, err := json.Marshal(config)
	if err != nil {
		log.Printf("Failed to marshal config: %v", err)
		return
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Failed to call endpoint: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Init client request to %s returned status %s", endpoint, resp.Status)

}
