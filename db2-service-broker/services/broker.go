package services

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"

	"github.com/astaxie/beego/logs"
)

var (
	db2ServiceBroker *DB2ServiceBroker
	lastOperationMap    map[string]int64
)

type DB2ServiceBroker struct {
}

func DB2ServiceBrokerInstance() *DB2ServiceBroker {
	if db2ServiceBroker == nil {
		db2ServiceBroker = &DB2ServiceBroker{}
	}
	return db2ServiceBroker
}

func (o *DB2ServiceBroker) Catalog() *brokerapi.Catalog {
	return readBrokerSettings()
}

func (o *DB2ServiceBroker) ServiceInstance(id string) (userProvidedServiceInstance, error) {
	result := userProvidedServiceInstance{}

	client := GetEtcdClientInstance()
	if client == nil {
		return result, errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + id)
	if err != nil {
		return result, errors.New("Get instance failre. The instance id is " + id)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	ok := true
	if serviceInstance.Name == "" {
		ok = false
	}
	if ok {
		result = serviceInstance
	}

	return result, nil
}

func (o *DB2ServiceBroker) Provision(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {

	DashboardURL := ""

	credential, ok := req.Parameters["credentials"]

	if !ok {
		return nil, errors.New("Parameters need to be provided \\'credential\\'")
	} else {
		jsonCred, err := json.Marshal(credential)
		if err != nil {
			glog.Errorf("Failed to marshal credentials: %v", err)
			return nil, err
		}
		cred := &brokerapi.Credential{}
		err = json.Unmarshal(jsonCred, cred)

		connectURI := (*cred)["connect_uri"]
		serviceName := (*cred)["service_name"]
		planName := (*cred)["plan_name"]
		if connectURI == "" {
			return nil, errors.New("Parameters need to be provided \\'connect_uri\\'")
		}
		if serviceName == "" {
			return nil, errors.New("Parameters need to be provided \\'service_name\\'")
		}
		if planName == "" {
			return nil, errors.New("Parameters need to be provided \\'plan_name\\'")
		}
		logs.Info("Parameters \"connect_uri\" : ", connectURI)
		logs.Info("Parameters \"service_name\" : ", serviceName)
		logs.Info("Parameters \"plan_name\" : ", planName)


		plan := getEqualPlan(serviceName.(string), planName.(string))
		if plan == nil {
			return nil, errors.New("Plan not found.Please select corrected plan.")
		}
		planValue := getPlanValue(plan)
		if planValue == "" {
			logs.Info("Get plan value equals \"\".")
			logs.Info("Get plan value :" , planValue)
			return nil, errors.New("Get plan value equals \"\".")
		}
		bufferPoolName, tablespaceName, userName, userPassword, err := createDatabaseAndUser(connectURI.(string), planValue)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("CRUD - Create database and user error.")
		}
		logs.Info("bufferPoolName: ", *bufferPoolName)
		logs.Info("tablespaceName: ", *tablespaceName)
		logs.Info("userName: ", *userName)
		logs.Info("userPassword: ", *userPassword)

		DashboardURLPointer, err := generateDB2Uri(connectURI.(string))
		if err != nil {
			logs.Error(err)
			return nil, errors.New("CRUD - Generate database dashboard url error.")
		}
		DashboardURL = *DashboardURLPointer
		(*cred)["gen_bufferpool"] = *bufferPoolName
		(*cred)["gen_tablespace"] = *tablespaceName
		(*cred)["gen_username"] = *userName
		(*cred)["gen_password"] = *userPassword

		//o.instanceMap[id] = &userProvidedServiceInstance{
		//	Name:       id,
		//	Credential: cred,
		//}

		contents, err := json.Marshal(userProvidedServiceInstance{
			Name:       id,
			Credential: cred,
		})
		if err != nil {
			logs.Info("Instance info switch failue.")
			return nil, errors.New("Instance info switch failue.")
		}
		client := GetEtcdClientInstance()
		if client == nil {
			logs.Info("Create etcd client instance failure.")
			return nil, errors.New("Create etcd client instance failure.")
		}
		logs.Info(string(contents))
		client.Set("/serviceinstance/"+id, string(contents))
	}

	//glog.Info("instance map len :", len(o.instanceMap))
	//glog.Info("instance map :", o.instanceMap)

	return &brokerapi.CreateServiceInstanceResponse{
		Operation:    "Provision",
		DashboardURL: DashboardURL,
	}, nil
}

func (o *DB2ServiceBroker) DeProvision(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	//o.rwMutex.Lock()
	//defer o.rwMutex.Unlock()

	//TODO: Need to be replace by etcd v3.
	//instance, ok := o.instanceMap[id]
	client := GetEtcdClientInstance()
	if client == nil {
		return nil, errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + id)
	if err != nil {
		return nil, errors.New("Get instance failre. The instance id is " + id)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	ok := true
	if serviceInstance.Name == "" {
		ok = false
	}

	if ok {
		cred := serviceInstance.Credential
		connectURI := (*cred)["connect_uri"]
		if connectURI == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'connect_uri\\'")
		}
		bufferpool := (*cred)["gen_bufferpool"]
		if bufferpool == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'gen_bufferpool\\'")
		}
		tablespace := (*cred)["gen_tablespace"]
		if tablespace == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'gen_tablespace\\'")
		}
		userName := (*cred)["gen_username"]
		if userName == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'gen_username\\'")
		}

		err := deleteDatabaseAndUser(connectURI.(string), bufferpool.(string), tablespace.(string), userName.(string))
		if err != nil {
			return nil, errors.New("CRUD - Delete database and user error.")
		}

		//delete(o.instanceMap, id)
		client.Delete("/serviceinstance/" + id)

		return &brokerapi.DeleteServiceInstanceResponse{
			Operation: "DeProvision",
		}, nil
	}

	return &brokerapi.DeleteServiceInstanceResponse{
		Operation: "DeProvision",
	}, nil
}

func (o *DB2ServiceBroker) Binding(instanceID, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	//o.rwMutex.RLock()
	//defer o.rwMutex.RUnlock()

	//TODO: Need to be replace by etcd v3.
	//instance, ok := o.instanceMap[instanceID]
	client := GetEtcdClientInstance()
	if client == nil {
		return nil, errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + instanceID)
	if err != nil {
		return nil, errors.New("Get instance failre. The instance id is " + instanceID)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	if serviceInstance.Name == "" {
		return nil, errors.New("no such instance: " + instanceID)
	}
	//if !ok {
	//	return nil, errors.New("no such instance: " + instanceID)
	//}
	//cred := instance.Credential

	// remove connect_uri from service instance credential.
	credential := serviceInstance.Credential
	delete(*credential, "connect_uri")
	serviceInstance.Credential = credential

	return &brokerapi.CreateServiceBindingResponse{
		Credentials: *(serviceInstance.Credential),
	}, nil
}

func (o *DB2ServiceBroker) UnBinding(instanceId, bindingId string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}

func (o *DB2ServiceBroker) ServiceInstanceLastOperation(instanceID, serviceID, planID, operation string) (*brokerapi.LastOperationResponse, error) {
	lower := strings.ToLower(operation)

	result := &brokerapi.LastOperationResponse{}

	client := GetEtcdClientInstance()
	if client == nil {
		return nil, errors.New("Create etcd client instance failure.")
	}
	Loop:
	switch lower {
	case "provision":
		_, err := client.Get("/serviceinstance/" + instanceID)

		if _, exist := lastOperationMap["provision"+instanceID]; exist {
			// key found.
			timeBefore := lastOperationMap["provision"+instanceID]
			duration := time.Now().Unix() - timeBefore
			if duration > 10 {
				if err != nil {
					result.State = "faild"
					result.Description = "Provision with instance " + instanceID + " faild)."
				} else {
					result.State = "succeeded"
					result.Description = "Provision with instance " + instanceID + " successed."
				}
				delete(lastOperationMap, "provision"+instanceID)
				break Loop
			}
		}

		if err != nil {
			result.State = "in progress"
			result.Description = "Provision with instance " + instanceID + " (10% complete)."
			if _, exist := lastOperationMap["provision"+instanceID]; !exist {
				lastOperationMap["provision"+instanceID] = time.Now().Unix()
			}
		} else {
			result.State = "succeeded"
			result.Description = "Provision with instance " + instanceID + " successed."
		}
	case "deprovision":
		_, err := client.Get("/serviceinstance/" + instanceID)
		if _, exist := lastOperationMap["deprovision"+instanceID]; exist {
			// key found.
			timeBefore := lastOperationMap["deprovision"+instanceID]
			duration := time.Now().Unix() - timeBefore
			if duration > 10 {
				if err != nil {
					result.State = "succeeded"
					result.Description = "DeProvision with instance " + instanceID + " successed."
				} else {
					result.State = "faild"
					result.Description = "DeProvision with instance " + instanceID + " faild."
				}
				delete(lastOperationMap, "deprovision"+instanceID)
				break Loop
			}
		}

		if err != nil {
			result.State = "succeeded"
			result.Description = "DeProvision with instance " + instanceID + " successed."
		} else {
			result.State = "in progress"
			result.Description = "DeProvision with instance " + instanceID + " (10% complete)."
			if _, exist := lastOperationMap["deprovision"+instanceID]; !exist {
				lastOperationMap["deprovision"+instanceID] = time.Now().Unix()
			}
		}
	case "update":
		result.State = "succeeded"
		result.Description = "Update with instance " + instanceID + " successed."
	default:
		return nil, errors.New("Unimplemented")
	}

	return result, nil
}