# Property Listing API

A production-ready REST API built with Beego (Go) that aggregates property listings from external services.

## Features

- ✅ RESTful API design with clean architecture
- ✅ API key authentication
- ✅ Concurrent property fetching with configurable limits
- ✅ Proper error handling and validation
- ✅ Request/response transformation
- ✅ Configuration management
- ✅ Structured logging
- ✅ Production-ready code structure

## Project Structure

```
property-listing-api/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── conf/
│   └── app.conf              # Application configuration
├── controllers/
│   └── property_controller.go # Request handlers
├── middleware/
│   └── auth.go               # Authentication middleware
├── models/
│   └── models.go             # Data models and structures
├── routers/
│   └── router.go             # Route definitions
└── services/
    ├── external_api.go       # External API client
    └── transformer.go        # Data transformation logic
```

## Prerequisites

- Go 1.21 or higher
- Access to external property services

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd property-listing-api
```

2. Install dependencies:
```bash
go mod download
```

3. Update configuration in `conf/app.conf`:
```ini
api_key = your_secret_api_key
location_service_url = http://192.168.0.35:8099/api
property_service_url = http://192.168.0.35:8099/api
```

## Running the Application

### Development Mode
```bash
go run main.go
```

### Production Mode
```bash
# Build
go build -o property-listing-api

# Run
./property-listing-api
```

The API will start on `http://localhost:8080`

## API Documentation

### Get Properties

Fetches property listings for a specific location.

**Endpoint:** `GET /v1/properties/:location`

**Headers:**
```
x-api-key: 63f4945d921d599f27ae4fdf5bada3f1
```

**Query Parameters:**
- `items` (required): Must be `true`

**URL Parameters:**
- `location`: Colon-separated location (e.g., `usa:florida:destin`)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/v1/properties/usa:florida:destin?items=true" \
  -H "x-api-key: 63f4945d921d599f27ae4fdf5bada3f1"
```

**Success Response (200):**
```json
{
  "Items": [
    {
      "ID": "BC-12810439",
      "Feed": 11,
      "Published": true,
      "GeoInfo": {
        "Categories": [...],
        "City": "Nārāyangarh",
        "Country": "Central Development Region",
        "CountryCode": "NP",
        "Display": "Bharatpur, Nepal",
        "LocationID": "571",
        "Lat": "27.627859",
        "Lng": "84.40818",
        "Slug": "nepal/bharatpur"
      },
      "Property": {
        "Amenities": {...},
        "Counts": {
          "Bedroom": 3,
          "Bathroom": 3,
          "Occupancy": 3
        },
        "FeatureImage": "...",
        "IsPetFriendly": true,
        "MinStay": 1,
        "PropertyName": "Dhakal Villa",
        "PropertySlug": "dhakal-villa",
        "PropertyType": "Casa",
        "PropertyTypeCategoryId": "6",
        "RoomSize": 1341.9
      },
      "Partner": {
        "ID": "12810439",
        "OwnerID": "",
        "Archived": [],
        "PropertyType": "Homestays",
        "URL": "https://www.booking.com/hotel/..."
      }
    }
  ]
}
```

**Error Responses:**

400 Bad Request:
```json
{
  "error": "location parameter is required"
}
```

401 Unauthorized:
```json
{
  "error": "x-api-key header is required"
}
```

502 Bad Gateway:
```json
{
  "error": "failed to fetch property IDs from location service"
}
```

500 Internal Server Error:
```json
{
  "error": "unexpected server error"
}
```

## Configuration

Edit `conf/app.conf` to customize:

- `httpport`: Server port (default: 8080)
- `api_key`: Secret API key for authentication
- `location_service_url`: External location service URL
- `property_service_url`: External property details service URL
- `external_api_timeout`: Timeout for external API calls (seconds)
- `max_concurrent_requests`: Max concurrent property detail requests

## Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./controllers -v
```

## Performance Optimization

The API uses concurrent goroutines to fetch property details in parallel:
- Configurable concurrency limit via `max_concurrent_requests`
- Semaphore pattern to prevent overwhelming external services
- Results maintain original order from location service

## Error Handling

The API implements comprehensive error handling:
- **Validation errors (400)**: Invalid parameters
- **Authentication errors (401)**: Missing or invalid API key
- **External service errors (502)**: Upstream service failures
- **Server errors (500)**: Internal processing errors

Failed property fetches are logged but don't fail the entire request.

## Security

- API key authentication via headers
- Input validation on all parameters
- No sensitive data in error messages
- Configurable timeouts to prevent resource exhaustion
