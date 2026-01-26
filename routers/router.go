package routers

import (
	"property_listing_api/controllers"
	"property_listing_api/middleware"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// API v1 namespace
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.APIKeyAuth),
		beego.NSNamespace("/properties",
			beego.NSRouter("/", &controllers.PropertyController{}, "get:GetProperties"),
			beego.NSRouter("/:location", &controllers.PropertyController{}, "get:GetProperties"),
		),
	)
	
	beego.AddNamespace(ns)
}
