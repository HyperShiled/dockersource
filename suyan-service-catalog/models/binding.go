package models

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
)

type PutBinding struct {
	Spec BindingSpec `json:"spec"`
}

type BindingMetadata struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

type BindingSpec struct {
	SecretName string `json:"secret_name"`
}

type BindingSpecInstanceRef struct {
	Name string `json:"name"`
}

type GetBinding struct {
	Binding *v1alpha1.Binding
}

type DeleteResultBinding struct {
	Message string `json:"message"`
}