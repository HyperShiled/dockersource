package services

import (
	"time"
	"errors"
	"suyan-service-catalog/models"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"github.com/astaxie/beego/logs"
)

type InstanceService struct {

}

func (s *InstanceService) Create(name string, namespace string, putInstance *models.PutInstance) (*models.GetInstance, error)  {
	logs.Info(name)
	logs.Info(namespace)
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	instanceList, err := clientSet.ServicecatalogV1alpha1().Instances(v1.NamespaceAll).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(instanceList.Items); i++ {
		instance := instanceList.Items[i]
		if instance.Name == name {
			return nil, errors.New("This instance already exists.")
		}
	}
	labels := make(map[string]string)
	labels[INSTANCE_LABELS_SERVICE_CLASS_NAME] = putInstance.Spec.ServiceClassName
	instance := &v1alpha1.Instance{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
			Namespace: namespace,
			Labels: labels,
		},
		Spec: v1alpha1.InstanceSpec{
			ServiceClassName: putInstance.Spec.ServiceClassName,
			PlanName: putInstance.Spec.PlanName,
			Parameters: putInstance.Spec.Parameters,
		},
	}
	instance, err = clientSet.ServicecatalogV1alpha1().Instances(namespace).Create(instance)
	if err != nil {
		return nil, err
	}
	instanceWatch, err := clientSet.ServicecatalogV1alpha1().Instances(namespace).Watch(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	InstanceLoop:
	for {
		select {
		case data := <-instanceWatch.ResultChan():
			broker := data.Object.(*v1alpha1.Instance)
			if data.Type == "ADDED" && broker.Name == name {
				break InstanceLoop
			}
		case <-time.After(time.Duration(10) * time.Second):
			break InstanceLoop
		}
	}
	instance, err = clientSet.ServicecatalogV1alpha1().Instances(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if instance != nil {
		return &models.GetInstance{
			Instance: instance,
		}, nil
	}
	return nil, nil
}

func (s *InstanceService) Update(namespace string, name string, patchInstance *models.PatchInstance) (*models.GetInstance, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	instance := &v1alpha1.Instance{
		Spec: v1alpha1.InstanceSpec{
			ServiceClassName: patchInstance.Spec.ServiceClassName,
			PlanName: patchInstance.Spec.PlanName,
		},
	}
	instance, err = clientSet.ServicecatalogV1alpha1().Instances(namespace).Update(instance)
	if err != nil {
		return nil, err
	}
	instance, err = clientSet.ServicecatalogV1alpha1().Instances(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if instance != nil {
		return &models.GetInstance{
			Instance: instance,
		}, nil
	}
	return nil, nil
}

func (s *InstanceService) Delete(namespace string, name string) (models.DeleteResultInstance, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return models.DeleteResultInstance{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}
	labelSelector := labels.SelectorFromSet(labels.Set{
		BINDING_LABELS_INSTANCE_NAME: name,
	})
	bindings, err := clientSet.ServicecatalogV1alpha1().Bindings("").List(v1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return models.DeleteResultInstance{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}
	if len(bindings.Items) > 0 {
		return models.DeleteResultInstance{
			Status: models.OPERATE_RESULT_FAILURE,
			Message: "There is a binding relationship under the current instance, cannot be deleted.",
		}, nil
	}
	err = clientSet.ServicecatalogV1alpha1().Instances(namespace).Delete(name, &v1.DeleteOptions{})
	if err != nil {
		return models.DeleteResultInstance{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}
	return models.DeleteResultInstance{
		Status: models.OPERATE_RESULT_SUCCESS,
		Message: "service instance delete success.",
	}, nil
}

func (s *InstanceService) Get(namespace string, name string) (*models.GetInstance, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	instance, err := clientSet.ServicecatalogV1alpha1().Instances(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if instance != nil {
		return &models.GetInstance{
			Instance: instance,
		}, nil
	}
	return nil, nil
}

func (s *InstanceService) GetDetail(namespace string, name string) (*models.GetInstanceItem, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	instance, err := clientSet.ServicecatalogV1alpha1().Instances(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// get binding list.
	labelSelector := labels.SelectorFromSet(labels.Set{
		BINDING_LABELS_INSTANCE_NAME: instance.Name,
	})
	bindingList, err := clientSet.ServicecatalogV1alpha1().Bindings(v1.NamespaceAll).List(v1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return nil, err
	}
	if instance != nil {
		return &models.GetInstanceItem{
			Instance: *instance,
			BindingList: bindingList.Items,
		}, nil
	}
	return nil, nil
}

func (s *InstanceService) List(namespace string) (*models.GetInstanceList, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	if namespace == "all" {
		namespace = v1.NamespaceAll
	}
	instanceList, err := clientSet.ServicecatalogV1alpha1().Instances(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if instanceList != nil {
		getInstanceList := &models.GetInstanceList{
			Items: make([]models.GetInstanceItem, 0),
		}
		for _, instance := range instanceList.Items {
			labelSelector := labels.SelectorFromSet(labels.Set{
				BINDING_LABELS_INSTANCE_NAME: instance.Name,
			})
			bindingList, err := clientSet.ServicecatalogV1alpha1().Bindings(v1.NamespaceAll).List(v1.ListOptions{
				LabelSelector: labelSelector.String(),
			})
			if err != nil {
				return nil, err
			}
			getInstanceItem := models.GetInstanceItem{
				Instance: instance,
				BindingList: bindingList.Items,
			}
			getInstanceList.Items = append(getInstanceList.Items, getInstanceItem)
		}
		return getInstanceList, nil
	}
	return nil, nil
}

