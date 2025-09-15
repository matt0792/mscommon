package configreader

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/matt0792/mscommon/commonmodels"
)

func GetMSConfig() (commonmodels.MicroserviceConfig, error) {
	if err := godotenv.Load(); err != nil {
		return commonmodels.MicroserviceConfig{}, err
	}

	serviceLocation := os.Getenv("SERVICE_LOCATION")
	serviceId := os.Getenv("SERVICE_ID")
	serviceName := os.Getenv("SERVICE_NAME")

	return commonmodels.MicroserviceConfig{
		ServiceLocation: serviceLocation,
		ServiceID:       serviceId,
		ServiceName:     serviceName,
	}, nil
}

func GetInternalConfig() {}
