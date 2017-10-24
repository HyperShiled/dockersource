package controllers

import (
	"net/http"
	"suyan-service-catalog/services"
	"github.com/astaxie/beego"
)

// Operations about Catalog
type CatalogController struct {
	beego.Controller
}

// @Title Get Catalog List.
// @Description get catalog list.
// @Success 200 {object} models.GetCatalogList
// @Failure 500 The request processing failure.
// @router / [get]
func (c *CatalogController) GetList() {
	catalogService := new(services.Catalog)
	catalogList, err := catalogService.List()
	if err != nil {
		c.Data["json"] = CreateErrorData(err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		c.Data["json"] = catalogList
	}
	c.ServeJSON()
}
