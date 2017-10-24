package services

import (
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/astaxie/beego"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
)

const (
	INSTANCE_LABELS_SERVICE_CLASS_NAME = "reference-service-class-name"
	BINDING_LABELS_INSTANCE_NAME = "reference-instance-name"
)

func CreateKubernetesClientSet() (*clientset.Clientset, error) {
	// load context config.
	section, _ := beego.AppConfig.GetSection("kubernetes")
	clusterName := section["cluster_name"]
	clusterServer := section["cluster_server"]
	contextName := section["context_name"]
	contextCluster := section["context_cluster"]
	// create kubernetes context.
	config := clientcmdapi.NewConfig()
	config.Clusters[clusterName] = &clientcmdapi.Cluster{
		Server: clusterServer,
	}
	config.Contexts[contextName] = &clientcmdapi.Context{
		Cluster: contextCluster,
	}
	config.CurrentContext = contextName
	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, contextName, &clientcmd.ConfigOverrides{}, nil)
	clientConfig, err := clientBuilder.ClientConfig()
	if err != nil {
		return nil, err
	}
	clientSet, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}


func GetBasicSecretInfoFromKubernetes() (string, string) {
	section, _ := beego.AppConfig.GetSection("kubernetes")
	basicSecretName := section["basic_secret_name"]
	basicSecretNamespace := section["basic_secret_namespace"]
	return basicSecretNamespace, basicSecretName
}
