package initclient

import (
	"fmt"
	"os"

	"github.com/matt0792/mscommon/commonmodels"
	"github.com/matt0792/mscommon/s2s"
)

func Init() error {
	token := os.Getenv("SERVICE_TOKEN")
	providerLoc := os.Getenv("PROVIDER_LOCATION")

	config := commonmodels.MicroserviceConfig{
		ServiceLocation: os.Getenv("SERVICE_LOCATION"),
		ServiceName:     os.Getenv("SERVICE_NAME"),
		ServiceID:       os.Getenv("SERVICE_ID"),
	}

	client, err := s2s.NewClient("ms-clientprovider", providerLoc, token)
	if err != nil {
		return err
	}

	var res string
	err = client.CallMethod("AddServiceConfig", config, &res)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
