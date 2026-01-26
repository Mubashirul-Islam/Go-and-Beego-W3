package controllers

import (
	"property_listing_api/models"
	"property_listing_api/services"
	"sync"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// PropertyController handles property-related requests
type PropertyController struct {
	beego.Controller
}

// job represents a property fetching task
type job struct {
	index      int
	propertyID string
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
	
	// Fetch property details for each ID using worker pool
	propertyItems := make([]models.PropertyItem, 0, len(propertyIDs))
	
	// Get max concurrent requests from config
	maxConcurrent, err := beego.AppConfig.Int("max_concurrent_requests")
	if err != nil {
		maxConcurrent = 10
	}
	
	// Create job channel and results storage
	jobs := make(chan job, len(propertyIDs))
	results := make([]*models.PropertyItem, len(propertyIDs))
	errors := make([]error, len(propertyIDs))
	
	// Start worker pool
	var wg sync.WaitGroup
	for w := 0; w < maxConcurrent; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			// Process jobs from channel
			for j := range jobs {
				// Fetch property details
				details, err := apiService.GetPropertyDetails(j.propertyID)
				if err != nil {
					logs.Error("Failed to fetch property %s: %v", j.propertyID, err)
					errors[j.index] = err
					continue
				}
				
				// Transform to internal format
				item, err := services.TransformPropertyDetails(details)
				if err != nil {
					logs.Error("Failed to transform property %s: %v", j.propertyID, err)
					errors[j.index] = err
					continue
				}
				
				results[j.index] = item
			}
		}()
	}
	
	// Send jobs to workers
	for i, propertyID := range propertyIDs {
		jobs <- job{index: i, propertyID: propertyID}
	}
	close(jobs)
	
	// Wait for all workers to complete
	wg.Wait()
	
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