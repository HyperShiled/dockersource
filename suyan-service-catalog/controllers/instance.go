package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"net/http"
	"suyan-service-catalog/models"
	"suyan-service-catalog/services"
)

// Operations about Instances.
type InstanceController struct {
	beego.Controller
}

// @Title Create Service Instance.
// @Description create service instance.
// @Param	instance_name		path 	string			true		"The key for service instance name."
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Param	body		body 	models.PutInstance	true		"body for service instance content"
// @Success 200 {object} models.GetInstance
// @Failure 400 The request processing failure.
// @router /:instance_name [put]
func (s *InstanceController) PutInstance() {
	name := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	if name == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			var putInstance models.PutInstance
			json.Unmarshal(s.Ctx.Input.RequestBody, &putInstance)
			instanceService := new(services.InstanceService)
			getInstance, err := instanceService.Create(name, namespace, &putInstance)
			if err != nil {
				s.Data["json"] = CreateErrorData(err)
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				s.Data["json"] = getInstance
			}
		}
	}
	s.ServeJSON()
}

// @Title Get Service Instance By Name.
// @Description get service instance by name.
// @Param	instance_name		path 	string	true		"The key for service instance name."
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Success 200 {object} models.GetInstance
// @Failure 400 The request processing failure.
// @router /:instance_name/last_operation [get]
func (s *InstanceController) Get() {
	name := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	if name == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			instanceService := new(services.InstanceService)
			instance, err := instanceService.Get(namespace, name)
			if err != nil {
				s.Data["json"] = CreateErrorData(err)
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				s.Data["json"] = instance
			}
		}
	}
	s.ServeJSON()
}

// @Title Get Service Instance Detail By Name.
// @Description get service instance detail by name.
// @Param	instance_name		path 	string	true		"The key for service instance name."
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Success 200 {object} models.GetInstanceItem
// @Failure 400 The request processing failure.
// @router /:instance_name [get]
func (s *InstanceController) GetDetail() {
	name := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	if name == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			instanceService := new(services.InstanceService)
			instance, err := instanceService.GetDetail(namespace, name)
			if err != nil {
				s.Data["json"] = CreateErrorData(err)
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				s.Data["json"] = instance
			}
		}
	}
	s.ServeJSON()
}

// @Title Get Service Instance List.
// @Description get service instance list.
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Success 200 {object} models.GetInstanceList
// @Failure 400 The request processing failure.
// @router / [get]
func (s *InstanceController) GetList() {
	namespace := s.GetString("namespace")
	if namespace == "" {
		s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		instanceService := new(services.InstanceService)
		brokerList, err := instanceService.List(namespace)
		if err != nil {
			s.Data["json"] = CreateErrorData(err)
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			s.Data["json"] = brokerList
		}
	}
	s.ServeJSON()
}

// @Title Delete Service Instance.
// @Description delete service instance.
// @Param	instance_name		path 	string	true		"The key for service instance name."
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Success 200 {object} models.DeleteResultInstance
// @Failure 400 The request processing failure.
// @router /:instance_name [delete]
func (s *InstanceController) Delete() {
	name := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	if name == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			instanceService := new(services.InstanceService)
			result, err := instanceService.Delete(namespace, name)
			if err != nil {
				s.Data["json"] = CreateErrorData(err)
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				s.Data["json"] = result
			}
		}
	}
	s.ServeJSON()
}

// @Title Update Service Instance.
// @Description update service instance.
// @Param	instance_name		path 	string	true		"The key for service instance name."
// @Param	namespace		query 	string	true		"The key for service instance namespace."
// @Param	body			body 	models.PatchInstance	true		"body for instance content"
// @Success 200 {object} models.getBroker
// @Failure 400 The request processing failure.
// @router /:instance_name [patch]
func (s *InstanceController) Put() {
	name := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	if name == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			var patchInstance models.PatchInstance
			json.Unmarshal(s.Ctx.Input.RequestBody, &patchInstance)
			instanceService := new(services.InstanceService)
			getInstance, err := instanceService.Update(namespace, name, &patchInstance)
			if err != nil {
				s.Data["json"] = CreateErrorData(err)
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				s.Data["json"] = getInstance
			}
		}
	}
	s.ServeJSON()
}

// @Title Update Service Binding.
// @Description update service binding.
// @Param	instance_name		path 	string			true		"The key for broker name."
// @Param	namespace		query 	string			true		"The key for service instance and binding namespace."
// @Param	binding_name		path 	string			true		"The key for binding name."
// @Param	body			body 	models.PutBinding	true		"body for binding content"
// @Success 200 {object} models.GetBinding
// @Failure 400 The request processing failure.
// @router /:instance_name/service_bindings/:binding_name [put]
func (s *InstanceController) PutBindings() {
	instanceName := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	bindingName := s.GetString(":binding_name")
	if instanceName == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			if bindingName == "" {
				s.Data["json"] = CreateErrorData(errors.New("binding_name is empty."))
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				var putBinding models.PutBinding
				json.Unmarshal(s.Ctx.Input.RequestBody, &putBinding)
				bindingService := new(services.BindingService)
				getBinding, err := bindingService.Create(namespace, bindingName, instanceName, &putBinding)
				if err != nil {
					s.Data["json"] = CreateErrorData(err)
					s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				} else {
					s.Data["json"] = getBinding
				}
			}
		}
	}
	s.ServeJSON()
}

// @Title Delete Service Binding.
// @Description delete service binding.
// @Param	instance_name		path 	string			true		"The key for broker name."
// @Param	namespace		query 	string			true		"The key for service instance and binding namespace."
// @Param	binding_name		path 	string			true		"The key for binding name."
// @Success 200 {object} models.DeleteResultBinding
// @Failure 400 The request processing failure.
// @router /:instance_name/service_bindings/:binding_name [delete]
func (s *InstanceController) DeleteBindings() {
	instanceName := s.GetString(":instance_name")
	namespace := s.GetString("namespace")
	bindingName := s.GetString(":binding_name")
	if instanceName == "" {
		s.Data["json"] = CreateErrorData(errors.New("name is empty."))
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		if namespace == "" {
			s.Data["json"] = CreateErrorData(errors.New("namespace is empty."))
			s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			if bindingName == "" {
				s.Data["json"] = CreateErrorData(errors.New("binding_name is empty."))
				s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			} else {
				bindingService := new(services.BindingService)
				err := bindingService.Delete(namespace, bindingName, instanceName)
				if err != nil {
					s.Data["json"] = CreateErrorData(err)
					s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
				} else {
					s.Data["json"] = models.DeleteResultBinding{
						Message: "service binding delete success.",
					}
				}
			}
		}
	}
	s.ServeJSON()
}
