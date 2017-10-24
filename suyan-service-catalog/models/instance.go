package models

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

type PutInstance struct {
	Spec InstanceSpec `json:"spec"`
}

type InstanceMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type InstanceSpec struct {
	ServiceClassName string                `json:"service_class_name"`
	PlanName         string                `json:"plan_name"`
	Parameters       *runtime.RawExtension `json:"parameters,omitempty"`
}

type GetInstance struct {
	Instance *v1alpha1.Instance
}

type DeleteResultInstance struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PatchInstance struct {
	Spec InstanceSpec `json:"spec"`
}

type GetInstanceList struct {
	Items []GetInstanceItem `json:"instance_item"`
}

type GetInstanceItem struct {
	Instance    v1alpha1.Instance  `json:"instance"`
	BindingList []v1alpha1.Binding `json:"binding_list"`
}
