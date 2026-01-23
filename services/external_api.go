package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"property_listing_api/models"
	"time"
	
	beego "github.com/beego/beego/v2/server/web"
)

// ExternalAPIService handles communication with external APIs
type ExternalAPIService struct {
	client              *http.Client
	locationServiceURL  string
	propertyServiceURL  string
}

// NewExternalAPIService creates a new instance of ExternalAPIService
func NewExternalAPIService() (*ExternalAPIService, error) {
	timeout, err := beego.AppConfig.Int("external_api_timeout")
	if err != nil {
		timeout = 30
	}
	
	locationURL, err := beego.AppConfig.String("location_service_url")
	if err != nil {
		return nil, fmt.Errorf("location_service_url not configured")
	}
	
	propertyURL, err := beego.AppConfig.String("property_service_url")
	if err != nil {
		return nil, fmt.Errorf("property_service_url not configured")
	}
	
	return &ExternalAPIService{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		locationServiceURL: locationURL,
		propertyServiceURL: propertyURL,
	}, nil
}

// GetPropertyIDsByLocation fetches property IDs for a given location
func (s *ExternalAPIService) GetPropertyIDsByLocation(location string) ([]string, error) {
	url := fmt.Sprintf("%s/%s", s.locationServiceURL, location)
	
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch property IDs: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("location service returned status %d: %s", resp.StatusCode, string(body))
	}
	
	var propertyIDs []string
	if err := json.NewDecoder(resp.Body).Decode(&propertyIDs); err != nil {
		return nil, fmt.Errorf("failed to decode property IDs: %w", err)
	}
	
	return propertyIDs, nil
}

// GetPropertyDetails fetches details for a specific property
func (s *ExternalAPIService) GetPropertyDetails(propertyID string) (*models.PropertyDetailsResponse, error) {
	url := fmt.Sprintf("%s/%s", s.propertyServiceURL, propertyID)
	
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch property details for %s: %w", propertyID, err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("property service returned status %d for %s: %s", resp.StatusCode, propertyID, string(body))
	}
	
	var details models.PropertyDetailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("failed to decode property details for %s: %w", propertyID, err)
	}
	
	return &details, nil
}