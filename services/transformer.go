package services

import (
	"fmt"
	"property_listing_api/models"
	"strconv"
)

// TransformPropertyDetails transforms external API response to internal format
func TransformPropertyDetails(details *models.PropertyDetailsResponse) (*models.PropertyItem, error) {
	// Parse categories
	categories, err := models.ParseCategories(details.Categories)
	if err != nil {
		return nil, fmt.Errorf("failed to parse categories: %w", err)
	}
	
	// Transform categories
	transformedCategories := make([]models.Category, 0, len(categories))
	var geoSlug string
	
	for i, cat := range categories {
		transformedCategories = append(transformedCategories, models.Category{
			Name:    cat.Name,
			Slug:    cat.Slug,
			Type:    cat.Type,
			Display: cat.Display,
		})
		
		// Use the last (most specific) category slug for GeoInfo.Slug
		if i == len(categories)-1 {
			geoSlug = cat.Slug
		}
	}
	
	// Transform amenities
	amenities := make(map[string]string)
	for i, amenity := range details.AmenityCategories {
		amenities[strconv.Itoa(i+1)] = amenity
	}

	
	
	return &models.PropertyItem{
		ID:        details.ID,
		Feed:      details.Feed,
		Published: details.Published,
		GeoInfo: models.GeoInfo{
			Categories:  transformedCategories,
			City:        details.City,
			Country:     details.Country,
			CountryCode: details.CountryCode,
			Display:     details.Display,
			LocationID:  details.LocationID,
			Lng:         details.LonLat.Coordinates[0],
			Lat:         details.LonLat.Coordinates[1],
			Slug:        geoSlug,
		},
		Property: models.Property{
			Amenities: amenities,
			Counts: models.Counts{
				Bedroom:   details.BedroomCount,
				Bathroom:  details.BathroomCount,
				Occupancy: details.Occupancy,
			},
			FeatureImage:           details.FeatureImage,
			IsPetFriendly:          details.PropertyFlags.IsPetFriendly,
			PropertyName:           details.PropertyName,
			PropertySlug:           details.PropertySlug,
			PropertyType:           details.PropertyType,
			RoomSize:               details.RoomSizeSqft,
		},
		Partner: models.Partner{
			ID:           details.FeedProviderID,
			OwnerID:      details.OwnerID,
			Archived:     details.Archived,
			PropertyType: details.PropertyType,
			URL:          details.FeedProviderURL,
		},
	}, nil
}

// ValidateLocation validates the location format
func ValidateLocation(location string) error {
	if location == "" {
		return fmt.Errorf("location parameter is required")
	}
	
	return nil
}