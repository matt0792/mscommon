package commonmodels

type InternalConfig struct {
	MongoUri      string
	DBName        string
	ProviderLoc   string
	ServiceToken  string
	ServiceConfig MicroserviceConfig
}
