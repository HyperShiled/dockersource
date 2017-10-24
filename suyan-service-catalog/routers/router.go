// @APIVersion 1.0.0
// @Title Kubernetes Service Broker API
// @Description Kubernetes service broker has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"suyan-service-catalog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/brokers",
			beego.NSInclude(
				&controllers.BrokerController{},
			),
		),
		beego.NSNamespace("/catalogs",
			beego.NSInclude(
				&controllers.CatalogController{},
			),
		),
		beego.NSNamespace("/service_instances",
			beego.NSInclude(
				&controllers.InstanceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
