package commonmodels

type InternalConfig struct {
	MongoUri      string
	DBName        string
	ProviderLoc   string
	ServiceToken  string
	Port          string
	ServiceConfig MicroserviceConfig
}
