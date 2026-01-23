package models

import "encoding/json"

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// PropertyListResponse is the main API response
type PropertyListResponse struct {
	Items []PropertyItem `json:"Items"`
}

// PropertyItem represents a single property in the response
type PropertyItem struct {
	ID        string       `json:"ID"`
	Feed      int          `json:"Feed"`
	Published bool         `json:"Published"`
	GeoInfo   GeoInfo      `json:"GeoInfo"`
	Property  Property     `json:"Property"`
	Partner   Partner      `json:"Partner"`
}

// GeoInfo contains geographical information
type GeoInfo struct {
	Categories  []Category `json:"Categories"`
	City        string     `json:"City"`
	Country     string     `json:"Country"`
	CountryCode string     `json:"CountryCode"`
	Display     string     `json:"Display"`
	LocationID  string     `json:"LocationID"`
	Lat         float64     `json:"Lat"`
	Lng         float64     `json:"Lng"`
	Slug        string     `json:"Slug"`
}

// Category represents a location category
type Category struct {
	Name    string   `json:"Name"`
	Slug    string   `json:"Slug"`
	Type    string   `json:"Type"`
	Display []string `json:"Display"`
}

// Property contains property-specific details
type Property struct {
	Amenities              map[string]string `json:"Amenities"`
	Counts                 Counts            `json:"Counts"`
	FeatureImage           string            `json:"FeatureImage"`
	IsPetFriendly          bool              `json:"IsPetFriendly"`
	PropertyName           string            `json:"PropertyName"`
	PropertySlug           string            `json:"PropertySlug"`
	PropertyType           string            `json:"PropertyType"`
	RoomSize               float64           `json:"RoomSize"`
}

// Counts contains bedroom, bathroom, and occupancy counts
type Counts struct {
	Bedroom   int `json:"Bedroom"`
	Bathroom  int `json:"Bathroom"`
	Occupancy int `json:"Occupancy"`
}

// Partner contains partner-specific information
type Partner struct {
	ID           string        `json:"ID"`
	OwnerID      string        `json:"OwnerID"`
	Archived     []interface{} `json:"Archived"`
	PropertyType string        `json:"PropertyType"`
	URL          string        `json:"URL"`
}

// External API Response Models

// PropertyDetailsResponse represents the response from property details API
type PropertyDetailsResponse struct {
	AmenityCategories       []string               `json:"amenity_categories"`
	Archived                []interface{}          `json:"archived"`
	BathroomCount           int                    `json:"bathroom_count"`
	BedroomCount            int                    `json:"bedroom_count"`
	Categories              string                 `json:"categories"`
	City                    string                 `json:"city"`
	Country                 string                 `json:"country"`
	CountryCode             string                 `json:"country_code"`
	Display                 string                 `json:"display"`
	FeatureImage            string                 `json:"feature_image"`
	Feed                    int                    `json:"feed"`
	FeedProviderID          string                 `json:"feed_provider_id"`
	FeedProviderURL         string                 `json:"feed_provider_url"`
	ID                      string                 `json:"id"`
	LocationID              string                 `json:"location_id"`
	LonLat                  LonLat                 `json:"lonlat"`
	Occupancy               int                    `json:"occupancy"`
	OwnerID                 string                 `json:"owner_id"`
	PropertyFlags           PropertyFlags          `json:"property_flags"`
	PropertyName            string                 `json:"property_name"`
	PropertySlug            string                 `json:"property_slug"`
	PropertyType            string                 `json:"property_type"`
	Published               bool                   `json:"published"`
	RoomSizeSqft            float64                `json:"room_size_sqft"`
}

// LonLat represents geographical coordinates
type LonLat struct {
	Coordinates []float64 `json:"coordinates"`
}

// PropertyFlags represents various property flags
type PropertyFlags struct {
	IsPetFriendly bool `json:"is_pet_friendly"`
}

// CategoryItem represents a parsed category item
type CategoryItem struct {
	LocationID string   `json:"LocationID"`
	Name       string   `json:"Name"`
	Type       string   `json:"Type"`
	Slug       string   `json:"Slug"`
	Display    []string `json:"Display"`
}

// Helper function to parse categories JSON string
func ParseCategories(categoriesJSON string) ([]CategoryItem, error) {
	var categories []CategoryItem
	if categoriesJSON == "" {
		return categories, nil
	}
	
	err := json.Unmarshal([]byte(categoriesJSON), &categories)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}