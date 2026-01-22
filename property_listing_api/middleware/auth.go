package middleware

import (
	"property_listing_api/models"
	
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// APIKeyAuth validates the API key from request header
func APIKeyAuth(ctx *context.Context) {
	apiKey := ctx.Input.Header("x-api-key")
	
	if apiKey == "" {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(models.ErrorResponse{
			Error: "x-api-key header is required",
		}, false, false)
		return
	}
	
	// Get expected API key from config
	expectedKey, err := beego.AppConfig.String("api_key")
	if err != nil {
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(models.ErrorResponse{
			Error: "internal server configuration error",
		}, false, false)
		return
	}
	
	if apiKey != expectedKey {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(models.ErrorResponse{
			Error: "invalid API key",
		}, false, false)
		return
	}
}