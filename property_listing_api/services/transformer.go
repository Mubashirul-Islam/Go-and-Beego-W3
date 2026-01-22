package services

import (
	"fmt"
	"property_listing_api/models"
	"strconv"
	"strings"
)

// Amenity translations mapping
var amenityTranslations = map[string]string{
	"Air Conditioner":      "Aire acondicionado",
	"Balcony/Terrace":      "Balcón/Terraza",
	"Hot Tub":              "Bañera de hidromasaje",
	"Internet":             "Internet",
	"Parking":              "Estacionamiento",
	"Pet Friendly":         "Mascota amigable",
	"View":                 "Vista",
	"Wellness Facilities":  "Instalaciones de bienestar",
	"Fireplace/Heating":    "Chimenea/Calefacción",
}

// Property type translations
var propertyTypeTranslations = map[string]string{
	"House":     "Casa",
	"Villa":     "Villa",
	"Apartment": "Apartamento",
	"Condo":     "Condominio",
	"Homestays": "Casa",
}

// Property type category IDs
var propertyTypeCategoryIDs = map[string]string{
	"House":     "6",
	"Villa":     "6",
	"Apartment": "1",
	"Condo":     "2",
}

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
	
	for _, cat := range categories {
		transformedCategories = append(transformedCategories, models.Category{
			Name:    cat.Name,
			Slug:    cat.Slug,
			Type:    cat.Type,
			Display: cat.Display,
		})
		
		// Use the last (most specific) category slug for GeoInfo.Slug
		if cat.Type == "city" || cat.Type == "region" {
			geoSlug = cat.Slug
		}
	}
	
	// Transform amenities
	amenities := make(map[string]string)
	for i, amenity := range details.AmenityCategories {
		if translated, ok := amenityTranslations[amenity]; ok {
			amenities[strconv.Itoa(i+1)] = translated
		} else {
			amenities[strconv.Itoa(i+1)] = amenity
		}
	}
	
	// Get coordinates
	lat := ""
	lng := ""
	if len(details.LonLat.Coordinates) >= 2 {
		lng = fmt.Sprintf("%f", details.LonLat.Coordinates[0])
		lat = fmt.Sprintf("%f", details.LonLat.Coordinates[1])
	}
	
	// Transform property type
	propertyType := details.PropertyTypeCategory
	if translated, ok := propertyTypeTranslations[propertyType]; ok {
		propertyType = translated
	}
	
	// Get property type category ID
	categoryID := "0"
	if id, ok := propertyTypeCategoryIDs[details.PropertyTypeCategory]; ok {
		categoryID = id
	}
	
	// Handle owner ID
	ownerID := ""
	if details.OwnerID != nil {
		ownerID = *details.OwnerID
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
			Lat:         lat,
			Lng:         lng,
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
			MinStay:                details.MinStay,
			PropertyName:           details.PropertyName,
			PropertySlug:           details.PropertySlug,
			PropertyType:           propertyType,
			PropertyTypeCategoryID: categoryID,
			RoomSize:               details.RoomSizeSqft,
		},
		Partner: models.Partner{
			ID:           details.FeedProviderID,
			OwnerID:      ownerID,
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
	
	parts := strings.Split(location, ":")
	if len(parts) < 2 {
		return fmt.Errorf("location must be colon-separated (e.g., usa:florida:destin)")
	}
	
	return nil
}