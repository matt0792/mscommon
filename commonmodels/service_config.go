package commonmodels

type MicroserviceConfigDocument struct {
	ID              string `bson:"_id,omitempty" json:"id"`
	ServiceLocation string `bson:"serviceLocation" json:"serviceLocation"`
	ServiceID       string `bson:"serviceId" json:"serviceId"`
	ServiceName     string `bson:"serviceName" json:"serviceName"`
}

type MicroserviceConfig struct {
	ServiceLocation string `bson:"serviceLocation" json:"serviceLocation"`
	ServiceID       string `bson:"serviceId" json:"serviceId"`
	ServiceName     string `bson:"serviceName" json:"serviceName"`
}
