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

func GetInternalConfig() (commonmodels.InternalConfig, error) {
	if err := godotenv.Load(); err != nil {
		return commonmodels.InternalConfig{}, err
	}

	serviceCfg, err := GetMSConfig()
	if err != nil {
		return commonmodels.InternalConfig{}, err
	}

	mongoUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	providerLoc := os.Getenv("PROVIDER_LOCATION")
	serviceToken := os.Getenv("SERVICE_TOKEN")

	return commonmodels.InternalConfig{
		MongoUri:      mongoUri,
		DBName:        dbName,
		ProviderLoc:   providerLoc,
		ServiceToken:  serviceToken,
		ServiceConfig: serviceCfg,
	}, nil
}
