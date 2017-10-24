package services

import (
	"time"
	"suyan-service-catalog/models"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	client_go_v1 "k8s.io/client-go/pkg/api/v1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	//"github.com/astaxie/beego/logs"
)

type BrokerService struct {

}

func (b *BrokerService) Create(postBroker *models.PostBroker) (*models.GetBroker, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	basicSecretNamespace, basicSecretName := GetBasicSecretInfoFromKubernetes()
	broker := &v1alpha1.Broker{
		ObjectMeta: v1.ObjectMeta{
			Name: postBroker.Metadata.Name,
		},
		Spec: v1alpha1.BrokerSpec{
			URL: postBroker.Spec.Url,
			AuthInfo: &v1alpha1.BrokerAuthInfo{
				BasicAuthSecret: &client_go_v1.ObjectReference{
					Namespace: basicSecretNamespace,
					Name: basicSecretName,
				},
			},
		},
	}
	broker, err = clientSet.ServicecatalogV1alpha1().Brokers().Create(broker)
	if err != nil {
		return nil, err
	}
	brokerWatch, err := clientSet.ServicecatalogV1alpha1().Brokers().Watch(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	BrokerLoop:
	for {
		select {
		case data := <-brokerWatch.ResultChan():
			broker := data.Object.(*v1alpha1.Broker)
			if data.Type == "MODIFIED" && broker.Name == postBroker.Metadata.Name && len(broker.Status.Conditions) > 0 {
				break BrokerLoop
			}
		case <-time.After(time.Duration(10) * time.Second):
			break BrokerLoop
		}
	}
	broker, err = clientSet.ServicecatalogV1alpha1().Brokers().Get(postBroker.Metadata.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if broker != nil {
		return &models.GetBroker{
			Broker: broker,
		}, nil
	}
	return nil, nil
}

func (b *BrokerService) Update(name string, putBroker *models.PutBroker) (*models.GetBroker, error) {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	broker, err := clientSet.ServicecatalogV1alpha1().Brokers().Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	broker.Spec.URL = putBroker.SpecUrl

	broker, err = clientSet.ServicecatalogV1alpha1().Brokers().Update(broker)
	if err != nil {
		return nil, err
	}
	if broker != nil {
		return &models.GetBroker{
			Broker: broker,
		}, nil
	}
	return nil, nil
}

func (b *BrokerService) Delete(name string) (models.DeleteResultBroker, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return models.DeleteResultBroker{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}

	serviceClasses, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().List(v1.ListOptions{})
	if err != nil {
		return models.DeleteResultBroker{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}
	serviceClassNameList := make([]string, 0)
	for _, serviceClass := range serviceClasses.Items {
		if serviceClass.BrokerName == name {
			serviceClassNameList = append(serviceClassNameList, serviceClass.Name)
		}
	}

	instanceNameList := make([]string, 0)
	for _, serviceClassName := range serviceClassNameList {
		labelSelector := labels.SelectorFromSet(labels.Set{
			INSTANCE_LABELS_SERVICE_CLASS_NAME: serviceClassName,
		})
		instanceList, err := clientSet.ServicecatalogV1alpha1().Instances("").List(v1.ListOptions{
			LabelSelector: labelSelector.String(),
		})
		if err != nil {
			return models.DeleteResultBroker{
				Status: models.OPERATE_RESULT_FAILURE,
			}, err
		}
		for _, instance := range instanceList.Items{
			instanceNameList = append(instanceNameList, instance.Name)
		}
	}
	if len(instanceNameList) > 0 {
		return models.DeleteResultBroker{
			Status: models.OPERATE_RESULT_FAILURE,
			Message: "Currently exists on the services provided by the broker service instance, unable to delete operation.",
		}, nil
	}

	err = clientSet.ServicecatalogV1alpha1().Brokers().Delete(name, &v1.DeleteOptions{})
	if err != nil {
		return models.DeleteResultBroker{
			Status: models.OPERATE_RESULT_FAILURE,
		}, err
	}
	return models.DeleteResultBroker{
		Status: models.OPERATE_RESULT_SUCCESS,
		Message: "broker delete success.",
	}, nil
}

func (b *BrokerService) Get(name string) (*models.GetBroker, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	broker, err := clientSet.ServicecatalogV1alpha1().Brokers().Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	serviceClassList, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var serviceClass v1alpha1.ServiceClass
	for i := 0; i < len(serviceClassList.Items); i++  {
		serviceClassItem := serviceClassList.Items[i]
		if serviceClassItem.BrokerName == broker.Name {
			serviceClass = serviceClassItem
			break
		}
	}
	getBroker := &models.GetBroker{
		Broker: broker,
		ServiceClass: &serviceClass,
	}
	if broker != nil {
		return getBroker, nil
	}
	return nil, nil
}

func (b *BrokerService) GetBrokerInstanceList(name string) (*models.GetInstanceList, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	broker, err := clientSet.ServicecatalogV1alpha1().Brokers().Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	serviceClassList, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var serviceClass v1alpha1.ServiceClass
	for i := 0; i < len(serviceClassList.Items); i++  {
		serviceClassItem := serviceClassList.Items[i]
		if serviceClassItem.BrokerName == broker.Name {
			serviceClass = serviceClassItem
			break
		}
	}
	labelSelector := labels.SelectorFromSet(labels.Set{
		INSTANCE_LABELS_SERVICE_CLASS_NAME: serviceClass.Name,
	})
	instanceList, err := clientSet.ServicecatalogV1alpha1().Instances(v1.NamespaceAll).List(v1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
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

func (b *BrokerService) List() (*models.GetBrokerList, error)  {
	clientSet, err := CreateKubernetesClientSet()
	if err != nil {
		return nil, err
	}
	brokerList, err := clientSet.ServicecatalogV1alpha1().Brokers().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	getBrokerList := models.GetBrokerList{
		BrokerList: make([]models.GetBroker, 0),
	}
	for i := 0; i < len(brokerList.Items); i++ {
		brokerItem := brokerList.Items[i]
		//logs.Debug("Broker Name: ", brokerItem.Name)
		getBroker := models.GetBroker{
			Broker: &brokerItem,
		}
		serviceClassList, err := clientSet.ServicecatalogV1alpha1().ServiceClasses().List(v1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for j := 0; j < len(serviceClassList.Items); j++  {
			serviceClassItem := serviceClassList.Items[j]
			if serviceClassItem.BrokerName == brokerItem.Name {
				getBroker.ServiceClass = &serviceClassItem
				break
			}
		}
		getBrokerList.BrokerList = append(getBrokerList.BrokerList, getBroker)
	}
	if err != nil {
		return nil, err
	}
	if brokerList != nil {
		return &getBrokerList, nil
	}
	return nil, nil
}
