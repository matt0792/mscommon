package initclient

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/matt0792/mscommon/commonmodels"
	"github.com/matt0792/mscommon/s2s"
)

func Init() {
	providerLoc := os.Getenv("PROVIDER_LOCATION")

	config := commonmodels.MicroserviceConfig{
		ServiceLocation: os.Getenv("SERVICE_LOCATION"),
		ServiceName:     os.Getenv("SERVICE_NAME"),
		ServiceID:       os.Getenv("SERVICE_ID"),
	}

	providerClient := s2s.NewClient(providerLoc, os.Getenv("SERVICE_TOKEN"))
	resp, err := providerClient.PostJSON("/services/register", config)
	if err != nil {
		log.Printf("Failed to call provider: %v", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response status: %s, body: %s", resp.Status, string(bodyBytes))
}
