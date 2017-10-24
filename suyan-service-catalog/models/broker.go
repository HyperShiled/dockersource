package models

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
)

type GetBrokerList struct {
	BrokerList []GetBroker
}

type GetBroker struct {
	Broker *v1alpha1.Broker
	ServiceClass *v1alpha1.ServiceClass
}

type PostBroker struct {
	Metadata BrokerMetadata `json:"metadata"`
	Spec     BrokerSpec     `json:"spec"`
}

type BrokerMetadata struct {
	Name string `json:"name"`
}

type BrokerSpec struct {
	Url string `json:"url"`
}

type DeleteResultBroker struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PutBroker struct {
	SpecUrl string `json:"spec_url"`
}
