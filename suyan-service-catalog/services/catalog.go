package services

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"suyan-service-catalog/models"
)

type Catalog struct {

}

func (c *Catalog) Create(serviceClass *v1alpha1.ServiceClass) (*v1alpha1.ServiceClass, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	serviceClass, err = clientSet.ServicecatalogV1alpha1().ServiceClasses().Create(serviceClass)
	if err != nil {
		return nil, err
	}
	return serviceClass, nil
}

func (c *Catalog) Update(serviceClass *v1alpha1.ServiceClass) (*v1alpha1.ServiceClass, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	serviceClass, err = clientSet.ServicecatalogV1alpha1().ServiceClasses().Update(serviceClass)
	if err != nil {
		return nil, err
	}
	return serviceClass, nil
}

func (c *Catalog) Delete(name string) (error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return err
	}
	err = clientSet.ServicecatalogV1alpha1().ServiceClasses().Delete(name, &v1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Catalog) Get(name string) (*v1alpha1.ServiceClass, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	serviceClass, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return serviceClass, nil
}

func (c *Catalog) List() (*models.GetCatalogList, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	serviceClassList, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if serviceClassList != nil {
		return &models.GetCatalogList{
			ServiceClassList: serviceClassList,
		}, nil
	}
	return nil, nil
}
