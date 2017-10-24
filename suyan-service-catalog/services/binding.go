package services

import (
	"time"
	"suyan-service-catalog/models"
	"k8s.io/client-go/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"errors"
)

type BindingService struct {

}

func (b *BindingService) Create(namespace string, name string, instanceName string, putBinding *models.PutBinding) (*models.GetBinding, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	labels := make(map[string]string)
	labels[BINDING_LABELS_INSTANCE_NAME] = instanceName
	binding := &v1alpha1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
			Labels: labels,
		},
		Spec: v1alpha1.BindingSpec{
			InstanceRef: v1.LocalObjectReference{
				Name: instanceName,
			},
			SecretName: putBinding.Spec.SecretName,
		},
	}
	binding, err = clientSet.ServicecatalogV1alpha1().Bindings(namespace).Create(binding)
	if err != nil {
		return nil, err
	}
	bindingWatch, err := clientSet.ServicecatalogV1alpha1().Bindings(namespace).Watch(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	BindingLoop:
	for {
		select {
		case data := <-bindingWatch.ResultChan():
			binding := data.Object.(*v1alpha1.Binding)
			if data.Type == "MODIFIED" && binding.Name == name {
				break BindingLoop
			}
		case <-time.After(time.Duration(10) * time.Second):
			break BindingLoop
		}
	}
	binding, err = clientSet.ServicecatalogV1alpha1().Bindings(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if binding != nil {
		return &models.GetBinding{
			Binding: binding,
		}, nil
	}
	return nil, nil
}

func (b *BindingService) Delete(namespace string, name string, instanceName string) (error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return err
	}
	binding, err := clientSet.ServicecatalogV1alpha1().Bindings(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if binding == nil {
		return errors.New("The binding does not exist.")
	}
	if binding.Spec.InstanceRef.Name != instanceName {
		return errors.New("The instance does not match.")
	}
	err = clientSet.ServicecatalogV1alpha1().Bindings(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
