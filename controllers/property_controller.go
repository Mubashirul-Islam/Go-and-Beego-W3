package controllers

import (
	"property_listing_api/models"
	"property_listing_api/services"
	"sync"
	
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

// PropertyController handles property-related requests
type PropertyController struct {
	beego.Controller
}

// GetProperties handles GET /v1/properties/:location
func (c *PropertyController) GetProperties() {
	// Get location from URL parameter
	location := c.Ctx.Input.Param(":location")
	
	// Validate location
	if err := services.ValidateLocation(location); err != nil {
		c.Data["json"] = models.ErrorResponse{Error: err.Error()}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	
	// Get items query parameter
	items := c.GetString("items")
	if items != "true" {
		c.Data["json"] = models.ErrorResponse{Error: "items parameter must be true"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	
	// Initialize external API service
	apiService, err := services.NewExternalAPIService()
	if err != nil {
		logs.Error("Failed to initialize API service: %v", err)
		c.Data["json"] = models.ErrorResponse{Error: "unexpected server error"}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}
	
	// Fetch property IDs
	propertyIDs, err := apiService.GetPropertyIDsByLocation(location)
	if err != nil {
		logs.Error("Failed to fetch property IDs: %v", err)
		c.Data["json"] = models.ErrorResponse{Error: "failed to fetch property IDs from location service"}
		c.Ctx.Output.SetStatus(502)
		c.ServeJSON()
		return
	}
	
	// Fetch property details for each ID
	propertyItems := make([]models.PropertyItem, 0, len(propertyIDs))
	
	// Get max concurrent requests from config
	maxConcurrent, err := beego.AppConfig.Int("max_concurrent_requests")
	if err != nil {
		maxConcurrent = 10
	}
	
	// Use goroutines with semaphore for concurrent requests
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrent)
	resultsChan := make(chan *models.PropertyItem, len(propertyIDs))
	errorsChan := make(chan error, len(propertyIDs))
	
	// Maintain order by using indices
	results := make([]*models.PropertyItem, len(propertyIDs))
	errors := make([]error, len(propertyIDs))
	
	for i, propertyID := range propertyIDs {
		wg.Add(1)
		go func(index int, id string) {
			defer wg.Done()
			
			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			
			// Fetch property details
			details, err := apiService.GetPropertyDetails(id)
			if err != nil {
				logs.Error("Failed to fetch property %s: %v", id, err)
				errors[index] = err
				return
			}
			
			// Transform to internal format
			item, err := services.TransformPropertyDetails(details)
			if err != nil {
				logs.Error("Failed to transform property %s: %v", id, err)
				errors[index] = err
				return
			}
			
			results[index] = item
		}(i, propertyID)
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	close(resultsChan)
	close(errorsChan)
	
	// Collect results in order, skipping failed ones
	for i, result := range results {
		if result != nil {
			propertyItems = append(propertyItems, *result)
		} else if errors[i] != nil {
			logs.Warn("Skipping property at index %d due to error: %v", i, errors[i])
		}
	}
	
	// Return response
	response := models.PropertyListResponse{
		Items: propertyItems,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}