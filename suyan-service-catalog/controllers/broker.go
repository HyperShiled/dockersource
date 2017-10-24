package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"net/http"
	"suyan-service-catalog/models"
	"suyan-service-catalog/services"
)

// Operations about Brokers
type BrokerController struct {
	beego.Controller
}

// @Title Get Broker List.
// @Description get broker list.
// @Success 200 {object} models.GetBrokerList
// @Failure 500 The request processing failure.
// @router / [get]
func (b *BrokerController) GetList() {
	brokerService := new(services.BrokerService)
	brokerList, err := brokerService.List()
	if err != nil {
		b.Data["json"] = CreateErrorData(err)
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		b.Data["json"] = brokerList
	}
	b.ServeJSON()
}

// @Title Get Broker By Name.
// @Description get broker by name.
// @Param	name		path 	string	true		"The key for broker name."
// @Success 200 {object} models.GetBroker
// @Failure 500 The request processing failure.
// @router /:name [get]
func (b *BrokerController) Get() {
	name := b.GetString(":name")
	if name == "" {
		b.Data["json"] = CreateErrorData(errors.New("name is empty."))
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		brokerService := new(services.BrokerService)
		broker, err := brokerService.Get(name)
		if err != nil {
			b.Data["json"] = CreateErrorData(err)
			b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			b.Data["json"] = broker
		}
	}
	b.ServeJSON()
}

// @Title Get Broker Service Instance List By Name.
// @Description get broker service instance list by name.
// @Param	name		path 	string	true		"The key for broker name."
// @Success 200 {object} models.GetBrokerInstanceList
// @Failure 500 The request processing failure.
// @router /:name/service_instances [get]
func (b *BrokerController) GetInstanceList() {
	name := b.GetString(":name")
	if name == "" {
		b.Data["json"] = CreateErrorData(errors.New("name is empty."))
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		brokerService := new(services.BrokerService)
		getBrokerInstanceList, err := brokerService.GetBrokerInstanceList(name)
		if err != nil {
			b.Data["json"] = CreateErrorData(err)
			b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			b.Data["json"] = getBrokerInstanceList
		}
	}
	b.ServeJSON()
}

// @Title Create Broker.
// @Description create broker.
// @Param	body		body 	models.PostBroker	true		"body for broker content"
// @Success 200 {object} models.GetBroker
// @Failure 500 The request processing failure.
// @router / [post]
func (b *BrokerController) Post() {
	var postBroker models.PostBroker
	json.Unmarshal(b.Ctx.Input.RequestBody, &postBroker)
	brokerService := new(services.BrokerService)
	getBroker, err := brokerService.Create(&postBroker)
	if err != nil {
		b.Data["json"] = CreateErrorData(err)
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		b.Data["json"] = getBroker
	}
	b.ServeJSON()
}

// @Title Delete Broker.
// @Description delete broker.
// @Param	name		path 	string	true		"The key for broker name."
// @Success 200 {object} models.DeleteResultBroker
// @Failure 500 The request processing failure.
// @router /:name [delete]
func (b *BrokerController) Delete() {
	name := b.GetString(":name")
	if name == "" {
		b.Data["json"] = CreateErrorData(errors.New("name is empty."))
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		brokerService := new(services.BrokerService)
		result, err := brokerService.Delete(name)
		if err != nil {
			b.Data["json"] = CreateErrorData(err)
			b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			b.Data["json"] = result
		}
	}
	b.ServeJSON()
}

// @Title Update Broker.
// @Description update broker.
// @Param	name		path 	string	true		"The key for broker name."
// @Param	body		body 	models.PutBroker	true		"body for broker content"
// @Success 200 {object} models.getBroker
// @Failure 500 The request processing failure.
// @router /:name [put]
func (b *BrokerController) Put() {
	name := b.GetString(":name")
	if name == "" {
		b.Data["json"] = CreateErrorData(errors.New("name is empty."))
		b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		var putBroker models.PutBroker
		json.Unmarshal(b.Ctx.Input.RequestBody, &putBroker)
		brokerService := new(services.BrokerService)
		getBroker, err := brokerService.Update(name, &putBroker)
		if err != nil {
			b.Data["json"] = CreateErrorData(err)
			b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		} else {
			b.Data["json"] = getBroker
		}
	}
	b.ServeJSON()
}
