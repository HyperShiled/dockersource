package models

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
)

type GetCatalogList struct {
	ServiceClassList *v1alpha1.ServiceClassList
}
