package commonmodels

type MicroserviceConfigDocument struct {
	ID              string `bson:"_id,omitempty" json:"id"`
	ServiceLocation string `bson:"service_location" json:"serviceLocation"`
	ServiceID       string `bson:"service_id" json:"serviceId"`
	ServiceName     string `bson:"service_name" json:"serviceName"`
}

type MicroserviceConfig struct {
	ServiceLocation string `bson:"service_location" json:"serviceLocation"`
	ServiceID       string `bson:"service_id" json:"serviceId"`
	ServiceName     string `bson:"service_name" json:"serviceName"`
}
