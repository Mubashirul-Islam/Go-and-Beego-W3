# Property Listing API

A production-ready REST API built with **Beego** (Go) that aggregates property listings from multiple external services with concurrent processing and robust error handling.

## Features

- **RESTful API Design** - Clean, predictable endpoints following REST principles
- **Secure Authentication** - API key-based authentication middleware
- **Concurrent Processing** - Parallel property fetching with configurable concurrency limits
- **Robust Error Handling** - Comprehensive validation and error responses
- **Data Transformation** - Request/response transformation pipeline
- **Configuration Management** - External configuration file support
- **Scalable Architecture** - Modular design with clear separation of concerns

---

## Project Structure

```
├── main.go                         # Application entry point
├── go.mod                          # Go module dependencies
├── go.sum                          # Dependency checksums
├── conf/
│   └── app.conf                    # Application configuration
├── controllers/
│   └── property_controller.go      # HTTP request handlers
├── middleware/
│   └── auth.go                     # Authentication middleware
├── models/
│   └── models.go                   # Data structures and models
├── routers/
│   └── router.go                   # API route definitions
└── services/
    ├── external_api.go             # External API integration
    └── transformer.go              # Data transformation logic
```

---

## 🔧 Prerequisites

- **Go**: Version 1.21 or higher
- **Network Access**: To external property and location services API
- **API Key**: Valid authentication credentials

---

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Mubashirul-Islam/Go-and-Beego-W3.git
cd Go-and-Beego-W3
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure the Application

Edit `conf/app.conf` with your settings:

```ini
# Server Configuration
appname = property-listing-api
httpport = 8080
runmode = dev

# Authentication
api_key = your_secret_api_key_here

# External Services
location_service_url = http://192.168.0.35:8099/api
property_service_url = http://192.168.0.35:8099/api

# Performance Settings
external_api_timeout = 30
max_concurrent_requests = 10
```

### 4.Start the Server

```bash
go run main.go
```

The server will start at `http://localhost:8080`


## API Documentation

### Endpoint: Get Properties

Retrieve property listings for a specific location with concurrent fetching from external services.

#### HTTP Request

```
GET /v1/properties/:location
```

#### Headers

| Parameter Type  | Name        | Data Type | Required | Description                                    | Value                              |
| --------------- | ----------- | --------- | -------- | ---------------------------------------------- | ---------------------------------- |
| Header          | `x-api-key` | string    | Yes      | API authentication key                         | `63f4945d921d599f27ae4fdf5bada3f1` |
| Path Variable   | `location`  | string    | Yes      | Colon-separated location hierarchy              | Example: `usa:texas`      |
| Query Parameter | `items`     | boolean   | Yes      | Flag to retrieve items                         | `true`                             |

#### Example Request

```bash
curl -X GET "http://localhost:8080/v1/properties/usa:florida:destin?items=true" \
  -H "x-api-key: 63f4945d921d599f27ae4fdf5bada3f1"
```

#### Success Response

**Status Code:** `200 OK`

<details>
  <summary>Click to expand</summary>

  
```json
{
    "Items": [
        {
            "ID": "HA-321898505",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fredericksburg",
                        "Slug": "usa/texas/fredericksburg",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fredericksburg"
                        ]
                    }
                ],
                "City": "Fredericksburg",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Fredericksburg, Texas, United States",
                "LocationID": "6059096",
                "Lat": 30.296688,
                "Lng": -98.873685,
                "Slug": "usa/texas/fredericksburg"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "Accessibility",
                    "12": "Security/Safety",
                    "13": "Sports/Activities",
                    "14": "Wellness Facilities",
                    "15": "Fireplace/Heating",
                    "16": "Guest Services",
                    "17": "Entertainment",
                    "18": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Internet",
                    "6": "Kitchen",
                    "7": "Laundry",
                    "8": "Parking",
                    "9": "Pet Friendly"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 3
                },
                "FeatureImage": "private-luxury-vi-us-fredericksburg-ha-321898505-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Private Luxury, Views, Nature, German History And Texas Hospitality!",
                "PropertySlug": "private-luxury-views-nature-german-history-and-texas-hospitality",
                "PropertyType": "Cottage",
                "RoomSize": 829
            },
            "Partner": {
                "ID": "898505",
                "OwnerID": "23711283",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Cottage",
                "URL": "https://www.vrbo.com/search?selected=23711283&regionId=6059096"
            }
        },
        {
            "ID": "BC-3983612",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Galveston",
                        "Slug": "usa/texas/galveston",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston"
                        ]
                    },
                    {
                        "Name": "San Jacinto",
                        "Slug": "usa/texas/galveston/san-jacinto",
                        "Type": "neighborhood",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston",
                            "san-jacinto"
                        ]
                    }
                ],
                "City": "Galveston",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Galveston, TX, USA",
                "LocationID": "553248635995267545",
                "Lat": 29.29925,
                "Lng": -94.77964,
                "Slug": "usa/texas/galveston/san-jacinto"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "2": "Internet",
                    "3": "Kitchen",
                    "4": "Laundry",
                    "5": "Parking"
                },
                "Counts": {
                    "Bedroom": 3,
                    "Bathroom": 2,
                    "Occupancy": 5
                },
                "FeatureImage": "cozy-historic-beach-cott-us-galveston-bc-3983612-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Cozy Historic Beach Cottage II",
                "PropertySlug": "cozy-historic-beach-cottage-ii",
                "PropertyType": "Holiday Homes",
                "RoomSize": 0
            },
            "Partner": {
                "ID": "3983612",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/cozy-tiny-house-two-blocks-from-the-beach.html"
            }
        },
        {
            "ID": "EP-91237705",
            "Feed": 24,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Houston",
                        "Slug": "usa/texas/houston",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "houston"
                        ]
                    },
                    {
                        "Name": "Astrodome",
                        "Slug": "usa/texas/houston/astrodome",
                        "Type": "neighborhood",
                        "Display": [
                            "usa",
                            "texas",
                            "houston",
                            "astrodome"
                        ]
                    }
                ],
                "City": "Houston",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Astrodome, Houston, Texas, USA",
                "LocationID": "553248635976007975",
                "Lat": 29.694173,
                "Lng": -95.406505,
                "Slug": "usa/texas/houston/astrodome"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "Wheelchair Accessible",
                    "12": "Accessibility",
                    "13": "Security/Safety",
                    "14": "Wellness Facilities",
                    "15": "Fireplace/Heating",
                    "16": "Guest Services",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Internet",
                    "6": "Kitchen",
                    "7": "Laundry",
                    "8": "Parking",
                    "9": "Pool"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 0
                },
                "FeatureImage": "packhouse-nrg-med-center-us-houston-ep-91237705-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "PackHouse NRG/MED Center",
                "PropertySlug": "packhouse-nrg-med-center",
                "PropertyType": "Apartment",
                "RoomSize": 0
            },
            "Partner": {
                "ID": "91237705",
                "OwnerID": "91237705",
                "Archived": [
                    "EP"
                ],
                "PropertyType": "Apartment",
                "URL": "https://www.expedia.com/Hotel-Search?regionId=553248635976007975&selected=91237705"
            }
        },
        {
            "ID": "HA-321454664",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fort Stockton",
                        "Slug": "usa/texas/fort-stockton",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fort-stockton"
                        ]
                    },
                    {
                        "Name": "Marfa",
                        "Slug": "usa/texas/marfa",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "marfa"
                        ]
                    }
                ],
                "City": "Alpine",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Marfa, Texas, USA",
                "LocationID": "55822",
                "Lat": 30.40588,
                "Lng": -103.683671,
                "Slug": "usa/texas/marfa"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Pet Friendly",
                    "11": "TV",
                    "12": "Security/Safety",
                    "13": "Wellness Facilities",
                    "14": "Fireplace/Heating",
                    "15": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 2
                },
                "FeatureImage": "rustic-western-experience-us-alpine-ha-321454664-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Rustic Western Experience With Hot Tub For Stargazing - No Extra Fees",
                "PropertySlug": "rustic-western-experience-with-hot-tub-for-stargazing-no-extra-fees",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 380
            },
            "Partner": {
                "ID": "454664",
                "OwnerID": "33498829",
                "Archived": [
                    "VRBO"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=33498829&regionId=6219254"
            }
        },
        {
            "ID": "BC-2255521",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Granbury",
                        "Slug": "usa/texas/granbury",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "granbury"
                        ]
                    },
                    {
                        "Name": "Glen Rose",
                        "Slug": "usa/texas/glen-rose",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "glen-rose"
                        ]
                    }
                ],
                "City": "Glen Rose",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "TX, USA",
                "LocationID": "108671",
                "Lat": 32.23843,
                "Lng": -97.753783,
                "Slug": "usa/texas/glen-rose"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Sports/Activities",
                    "11": "Wellness Facilities",
                    "12": "Spa",
                    "13": "Fireplace/Heating",
                    "14": "Restaurant",
                    "15": "Guest Services",
                    "16": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Child Friendly",
                    "4": "Hot Tub",
                    "5": "Internet",
                    "6": "Parking",
                    "7": "Designated Smoking Area",
                    "8": "Accessibility",
                    "9": "Security/Safety"
                },
                "Counts": {
                    "Bedroom": 5,
                    "Bathroom": 5,
                    "Occupancy": 16
                },
                "FeatureImage": "live-oak-b-b-us-glen-rose-bc-2255521-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Live Oak B&B",
                "PropertySlug": "live-oak-b-b",
                "PropertyType": "Bed And Breakfasts",
                "RoomSize": 232.5
            },
            "Partner": {
                "ID": "2255521",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Bed And Breakfasts",
                "URL": "https://www.booking.com/hotel/us/live-oak-b-b.html?aid=affiliate_id"
            }
        },
        {
            "ID": "EP-88731209",
            "Feed": 24,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Galveston",
                        "Slug": "usa/texas/galveston",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston"
                        ]
                    },
                    {
                        "Name": "Port Bolivar",
                        "Slug": "usa/texas/port-bolivar",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "port-bolivar"
                        ]
                    }
                ],
                "City": "Bolivar Peninsula",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Port Bolivar, Galveston, Texas, USA",
                "LocationID": "6141815",
                "Lat": 29.506823,
                "Lng": -94.502827,
                "Slug": "usa/texas/port-bolivar"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Ocean View",
                    "11": "Security/Safety",
                    "12": "Fireplace/Heating",
                    "2": "Balcony/Terrace",
                    "3": "Child Friendly",
                    "4": "Internet",
                    "5": "Kitchen",
                    "6": "Laundry",
                    "7": "Parking",
                    "8": "TV",
                    "9": "View"
                },
                "Counts": {
                    "Bedroom": 4,
                    "Bathroom": 2,
                    "Occupancy": 12
                },
                "FeatureImage": "lavender-rays-us-bolivar-peninsula-ep-88731209-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Lavender Rays",
                "PropertySlug": "lavender-rays",
                "PropertyType": "Apartment",
                "RoomSize": 1600
            },
            "Partner": {
                "ID": "88731209",
                "OwnerID": "88731209",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Apartment",
                "URL": "https://www.expedia.com/Hotel-Search?regionId=6141815&selected=88731209"
            }
        },
        {
            "ID": "HA-121138941",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "New Braunfels",
                        "Slug": "usa/texas/new-braunfels",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "new-braunfels"
                        ]
                    },
                    {
                        "Name": "Canyon Lake",
                        "Slug": "usa/texas/canyon-lake",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "canyon-lake"
                        ]
                    }
                ],
                "City": "Canyon Lake",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Canyon Lake, Texas, United States",
                "LocationID": "6050248",
                "Lat": 29.886641,
                "Lng": -98.277289,
                "Slug": "usa/texas/canyon-lake"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "View",
                    "12": "Ocean View",
                    "13": "Security/Safety",
                    "14": "Sports/Activities",
                    "15": "Wellness Facilities",
                    "16": "Fireplace/Heating",
                    "17": "Entertainment",
                    "18": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 3,
                    "Occupancy": 8
                },
                "FeatureImage": "modern-lakefront-hot-us-canyon-lake-ha-121138941-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Modern Lakefront, Hot tub, Firepits, Large multi-level decks, Amazing Sunsets",
                "PropertySlug": "modern-lakefront-hot-tub-firepits-large-multi-level-decks-amazing-sunsets",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 1300
            },
            "Partner": {
                "ID": "138941",
                "OwnerID": "31902003",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=31902003&regionId=6050248"
            }
        },
        {
            "ID": "BC-6597745",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Galveston",
                        "Slug": "usa/texas/galveston",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston"
                        ]
                    }
                ],
                "City": "Galveston",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Galveston, TX, USA",
                "LocationID": "1341",
                "Lat": 29.28388,
                "Lng": -94.80001,
                "Slug": "usa/texas/galveston"
            },
            "Property": {
                "Amenities": {
                    "1": "Internet",
                    "2": "Kitchen",
                    "3": "Parking"
                },
                "Counts": {
                    "Bedroom": 21,
                    "Bathroom": 1,
                    "Occupancy": 40
                },
                "FeatureImage": "ocean-house-hostel-us-galveston-bc-6597745-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Ocean House Hostel",
                "PropertySlug": "ocean-house-hostel",
                "PropertyType": "Hostels",
                "RoomSize": 0
            },
            "Partner": {
                "ID": "6597745",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Hostels",
                "URL": "https://www.booking.com/hotel/us/ocean-house-hostel.html"
            }
        },
        {
            "ID": "EP-2762080",
            "Feed": 24,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fort Worth",
                        "Slug": "usa/texas/fort-worth",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fort-worth"
                        ]
                    }
                ],
                "City": "Fort Worth",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Fort Worth, Texas, United States",
                "LocationID": "178258",
                "Lat": 32.73669,
                "Lng": -97.32934,
                "Slug": "usa/texas/fort-worth"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Wellness Facilities",
                    "11": "Toiletries",
                    "12": "Guest Services",
                    "13": "Entertainment",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Internet",
                    "5": "Laundry",
                    "6": "Parking",
                    "7": "Accessibility",
                    "8": "Security/Safety",
                    "9": "Business Services"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 0
                },
                "FeatureImage": "three-danes-inn-us-fort-worth-ep-2762080-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Three Danes Inn",
                "PropertySlug": "three-danes-inn",
                "PropertyType": "Guest House",
                "RoomSize": 0
            },
            "Partner": {
                "ID": "2762080",
                "OwnerID": "2762080",
                "Archived": [
                    "EP",
                    "HC"
                ],
                "PropertyType": "Guest House",
                "URL": "https://www.expedia.com/Hotel-Search?regionId=178258&selected=2762080"
            }
        },
        {
            "ID": "HA-3211013519",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Waco",
                        "Slug": "usa/texas/waco",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "waco"
                        ]
                    },
                    {
                        "Name": "Woodway",
                        "Slug": "usa/texas/waco/woodway",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "waco",
                            "woodway"
                        ]
                    }
                ],
                "City": "Waco",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Waco, Texas, United States",
                "LocationID": "6034031",
                "Lat": 31.477981,
                "Lng": -97.225179,
                "Slug": "usa/texas/waco/woodway"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "Security/Safety",
                    "12": "Sports/Activities",
                    "13": "Wellness Facilities",
                    "14": "Fireplace/Heating",
                    "15": "Entertainment",
                    "16": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Internet",
                    "6": "Kitchen",
                    "7": "Laundry",
                    "8": "Parking",
                    "9": "Pet Friendly"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 6
                },
                "FeatureImage": "the-cottage-chapel-ridge-1-us-waco-ha-3211013519-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "THE COTTAGE@CHAPEL RIDGE #1 in Waco for Exceptional reviews! Ranked #6 in all TX",
                "PropertySlug": "the-cottage-chapel-ridge-1-in-waco-for-exceptional-reviews-ranked-6-in-all-tx",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 1360
            },
            "Partner": {
                "ID": "1013519",
                "OwnerID": "19939139",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=19939139&regionId=6034031"
            }
        },
        {
            "ID": "BC-7876528",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "San Antonio",
                        "Slug": "usa/texas/san-antonio",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "san-antonio"
                        ]
                    },
                    {
                        "Name": "Near East Side",
                        "Slug": "usa/texas/san-antonio/near-east-side",
                        "Type": "neighborhood",
                        "Display": [
                            "usa",
                            "texas",
                            "san-antonio",
                            "near-east-side"
                        ]
                    }
                ],
                "City": "San Antonio",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "San Antonio, TX, USA",
                "LocationID": "553248635976482249",
                "Lat": 29.417081,
                "Lng": -98.470436,
                "Slug": "usa/texas/san-antonio/near-east-side"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Guest Services",
                    "11": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Breakfast",
                    "4": "Child Friendly",
                    "5": "Internet",
                    "6": "Parking",
                    "7": "Security/Safety",
                    "8": "Sports/Activities",
                    "9": "Fireplace/Heating"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 6
                },
                "FeatureImage": "spacious-townhouse-min-us-san-antonio-bc-7876528-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "SPACIOUS TOWNHOUSE MINUTES FROM DOWNTOWN SA.",
                "PropertySlug": "spacious-townhouse-minutes-from-downtown-sa",
                "PropertyType": "Holiday Homes",
                "RoomSize": 1194.79
            },
            "Partner": {
                "ID": "7876528",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/spacious-townhouse-minutes-from-downtown-sa.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-321196219",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "New Braunfels",
                        "Slug": "usa/texas/new-braunfels",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "new-braunfels"
                        ]
                    },
                    {
                        "Name": "Wimberley",
                        "Slug": "usa/texas/wimberley",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "wimberley"
                        ]
                    }
                ],
                "City": "Wimberley",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Wimberley, Texas, United States",
                "LocationID": "180937",
                "Lat": 29.979807,
                "Lng": -98.114563,
                "Slug": "usa/texas/wimberley"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Pet Friendly",
                    "11": "TV",
                    "12": "Security/Safety",
                    "13": "Sports/Activities",
                    "14": "Wellness Facilities",
                    "15": "Fireplace/Heating",
                    "16": "Guest Services",
                    "17": "Entertainment",
                    "18": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 1,
                    "Occupancy": 4
                },
                "FeatureImage": "charming-blanco-riverf-us-wimberley-ha-321196219-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Charming Blanco Riverfont Cottage",
                "PropertySlug": "charming-blanco-riverfont-cottage",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 1000
            },
            "Partner": {
                "ID": "196219",
                "OwnerID": "19007641",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=19007641&regionId=180937"
            }
        },
        {
            "ID": "BC-3156886",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Galveston",
                        "Slug": "usa/texas/galveston",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston"
                        ]
                    },
                    {
                        "Name": "Terramar Beach",
                        "Slug": "usa/texas/galveston/terramar-beach",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "galveston",
                            "terramar-beach"
                        ]
                    }
                ],
                "City": "Galveston",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Galveston, TX, USA",
                "LocationID": "3000045688",
                "Lat": 29.12991,
                "Lng": -95.06046,
                "Slug": "usa/texas/galveston/terramar-beach"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "2": "Balcony/Terrace",
                    "3": "Internet",
                    "4": "Kitchen",
                    "5": "Parking",
                    "6": "Pet Friendly"
                },
                "Counts": {
                    "Bedroom": 6,
                    "Bathroom": 2,
                    "Occupancy": 10
                },
                "FeatureImage": "blue-heron-house-us-galveston-bc-3156886-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Blue Heron House",
                "PropertySlug": "blue-heron-house",
                "PropertyType": "Holiday Homes",
                "RoomSize": 1440
            },
            "Partner": {
                "ID": "3156886",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/blue-heron-house.html"
            }
        },
        {
            "ID": "HA-321729987",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Waco",
                        "Slug": "usa/texas/waco",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "waco"
                        ]
                    },
                    {
                        "Name": "Woodway",
                        "Slug": "usa/texas/waco/woodway",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "waco",
                            "woodway"
                        ]
                    }
                ],
                "City": "Waco",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Waco, Texas, United States",
                "LocationID": "6034031",
                "Lat": 31.477889,
                "Lng": -97.225112,
                "Slug": "usa/texas/waco/woodway"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "Security/Safety",
                    "12": "Sports/Activities",
                    "13": "Wellness Facilities",
                    "14": "Fireplace/Heating",
                    "15": "Entertainment",
                    "16": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Breakfast",
                    "5": "Child Friendly",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 5
                },
                "FeatureImage": "the-haven-chapel-ridge-2-in-us-waco-ha-321729987-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "THE HAVEN @ CHAPEL RIDGE #2 in Waco for Exceptional Reviews! Ranked #7 in all TX",
                "PropertySlug": "the-haven-chapel-ridge-2-in-waco-for-exceptional-reviews-ranked-7-in-all-tx",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 1200
            },
            "Partner": {
                "ID": "729987",
                "OwnerID": "25795056",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=25795056&regionId=6034031"
            }
        },
        {
            "ID": "BC-4840589",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Waco",
                        "Slug": "usa/texas/waco",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "waco"
                        ]
                    }
                ],
                "City": "Waco",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Waco, TX, USA",
                "LocationID": "286",
                "Lat": 31.655746,
                "Lng": -97.143387,
                "Slug": "usa/texas/waco"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "2": "Internet",
                    "3": "Parking",
                    "4": "Security/Safety",
                    "5": "Fireplace/Heating"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 3
                },
                "FeatureImage": "cozy-cabin-little-red-hen-12-us-waco-bc-4840589-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Cozy Cabin Little Red Hen 12 min to Magnolia",
                "PropertySlug": "cozy-cabin-little-red-hen-12-min-to-magnolia",
                "PropertyType": "Holiday Homes",
                "RoomSize": 322.92
            },
            "Partner": {
                "ID": "4840589",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/cozy-cabin-little-red-hen-12-min-to-magnolia.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-321923603",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fort Stockton",
                        "Slug": "usa/texas/fort-stockton",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fort-stockton"
                        ]
                    },
                    {
                        "Name": "Marfa",
                        "Slug": "usa/texas/marfa",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "marfa"
                        ]
                    }
                ],
                "City": "Alpine",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Marfa, Texas, USA",
                "LocationID": "55822",
                "Lat": 30.405685,
                "Lng": -103.683463,
                "Slug": "usa/texas/marfa"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Pet Friendly",
                    "11": "TV",
                    "12": "Accessibility",
                    "13": "Security/Safety",
                    "14": "Wellness Facilities",
                    "15": "Fireplace/Heating",
                    "16": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 4
                },
                "FeatureImage": "sierra-vista-guest-house-us-alpine-ha-321923603-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Sierra Vista Guest House-With Hot Tub For Stargazing-No Extra Cleaning/Pet Fees",
                "PropertySlug": "sierra-vista-guest-house-with-hot-tub-for-stargazing-no-extra-cleaning-pet-fees",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 660
            },
            "Partner": {
                "ID": "923603",
                "OwnerID": "34274819",
                "Archived": [
                    "VRBO"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=34274819&regionId=6219254"
            }
        },
        {
            "ID": "BC-10164843",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fredericksburg",
                        "Slug": "usa/texas/fredericksburg",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fredericksburg"
                        ]
                    }
                ],
                "City": "Fredericksburg",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Fredericksburg, TX, USA",
                "LocationID": "6059096",
                "Lat": 30.237623,
                "Lng": -98.886438,
                "Slug": "usa/texas/fredericksburg"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "2": "Internet",
                    "3": "Parking",
                    "4": "Security/Safety",
                    "5": "Fireplace/Heating"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 2
                },
                "FeatureImage": "new-wilderness-hid-us-fredericksburg-bc-10164843-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "New Wilderness Hideout-Cozy Container Home",
                "PropertySlug": "new-wilderness-hideout-cozy-container-home",
                "PropertyType": "Holiday Homes",
                "RoomSize": 161.46
            },
            "Partner": {
                "ID": "10164843",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/new-wilderness-hideout-cozy-container-home.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-121358785",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "New Braunfels",
                        "Slug": "usa/texas/new-braunfels",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "new-braunfels"
                        ]
                    }
                ],
                "City": "Canyon Lake",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Canyon Lake, Texas, United States",
                "LocationID": "6059109",
                "Lat": 29.841157,
                "Lng": -98.171901,
                "Slug": "usa/texas/new-braunfels"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "TV",
                    "11": "Security/Safety",
                    "12": "Sports/Activities",
                    "13": "Wellness Facilities",
                    "14": "Fireplace/Heating",
                    "15": "Entertainment",
                    "16": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Internet",
                    "6": "Kitchen",
                    "7": "Laundry",
                    "8": "Parking",
                    "9": "Pet Friendly"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 6
                },
                "FeatureImage": "scenic-riverfront-ho-us-canyon-lake-ha-121358785-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Scenic riverfront home on quiet stretch of Guadalupe - waterfall by huge front porch",
                "PropertySlug": "scenic-riverfront-home-on-quiet-stretch-of-guadalupe-waterfall-by-huge-front-porch",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 1270
            },
            "Partner": {
                "ID": "358785",
                "OwnerID": "24724749",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=24724749&regionId=6059109"
            }
        },
        {
            "ID": "BC-4504976",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Austin",
                        "Slug": "usa/texas/austin",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "austin"
                        ]
                    },
                    {
                        "Name": "Pflugerville",
                        "Slug": "usa/texas/pflugerville",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "pflugerville"
                        ]
                    }
                ],
                "City": "Pflugerville",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Pflugerville, TX, USA",
                "LocationID": "6085432",
                "Lat": 30.459975,
                "Lng": -97.569408,
                "Slug": "usa/texas/pflugerville"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Ocean View",
                    "11": "Oceanfront",
                    "12": "Accessibility",
                    "13": "Security/Safety",
                    "14": "Sports/Activities",
                    "15": "Fireplace/Heating",
                    "16": "Guest Services",
                    "17": "Entertainment",
                    "2": "Balcony/Terrace",
                    "3": "Child Friendly",
                    "4": "Internet",
                    "5": "Parking",
                    "6": "Pool",
                    "7": "Designated Smoking Area",
                    "8": "TV",
                    "9": "View"
                },
                "Counts": {
                    "Bedroom": 4,
                    "Bathroom": 4,
                    "Occupancy": 12
                },
                "FeatureImage": "the-plantation-house-us-pflugerville-bc-4504976-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "The Plantation House Boutique Inn",
                "PropertySlug": "the-plantation-house-boutique-inn",
                "PropertyType": "Bed And Breakfasts",
                "RoomSize": 575.87
            },
            "Partner": {
                "ID": "4504976",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Bed And Breakfasts",
                "URL": "https://www.booking.com/hotel/us/the-plantation-house-boutique-inn.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-3212447553",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Decatur",
                        "Slug": "usa/texas/decatur",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "decatur"
                        ]
                    },
                    {
                        "Name": "Bridgeport",
                        "Slug": "usa/texas/bridgeport",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "bridgeport"
                        ]
                    }
                ],
                "City": "Bridgeport",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Bridgeport, Texas, USA",
                "LocationID": "6046718",
                "Lat": 33.172217,
                "Lng": -97.824742,
                "Slug": "usa/texas/bridgeport"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Pet Friendly",
                    "11": "Designated Smoking Area",
                    "12": "TV",
                    "13": "Sports/Activities",
                    "14": "Wellness Facilities",
                    "15": "Fireplace/Heating",
                    "16": "Entertainment",
                    "17": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 7
                },
                "FeatureImage": "open-modern-cabin-fi-us-bridgeport-ha-3212447553-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Open modern Cabin, fishing, shooting range, tranquil/private, firepit, 4W, canoe",
                "PropertySlug": "open-modern-cabin-fishing-shooting-range-tranquil-private-firepit-4w-canoe",
                "PropertyType": "Cabin",
                "RoomSize": 780
            },
            "Partner": {
                "ID": "2447553",
                "OwnerID": "70871952",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Cabin",
                "URL": "https://www.vrbo.com/search?selected=70871952&regionId=6046718"
            }
        },
        {
            "ID": "BC-10306427",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Bastrop",
                        "Slug": "usa/texas/bastrop",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "bastrop"
                        ]
                    }
                ],
                "City": "Smithville",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Bastrop, Texas, USA",
                "LocationID": "55572",
                "Lat": 30.047498,
                "Lng": -97.203271,
                "Slug": "usa/texas/bastrop"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Sports/Activities",
                    "11": "Wellness Facilities",
                    "12": "Spa",
                    "13": "Fireplace/Heating",
                    "14": "Guest Services",
                    "2": "Balcony/Terrace",
                    "3": "Child Friendly",
                    "4": "Internet",
                    "5": "Parking",
                    "6": "Pet Friendly",
                    "7": "Designated Smoking Area",
                    "8": "Security/Safety",
                    "9": "EV Charge Station"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 4
                },
                "FeatureImage": "river-rv-romance-glamp-us-smithville-bc-10306427-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "River RV, Romance Glamping Pup Paradise Riverside",
                "PropertySlug": "river-rv-romance-glamping-pup-paradise-riverside",
                "PropertyType": "Holiday Homes",
                "RoomSize": 247.57
            },
            "Partner": {
                "ID": "10306427",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/relax-by-the-fire-and-stargaze-by-the-river.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-3211234580",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fredericksburg",
                        "Slug": "usa/texas/fredericksburg",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fredericksburg"
                        ]
                    }
                ],
                "City": "Fredericksburg",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "Fredericksburg, Texas, United States",
                "LocationID": "6059096",
                "Lat": 30.382927,
                "Lng": -98.95792,
                "Slug": "usa/texas/fredericksburg"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Wellness Facilities",
                    "11": "Fireplace/Heating",
                    "12": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Internet",
                    "5": "Kitchen",
                    "6": "Parking",
                    "7": "Pet Friendly",
                    "8": "TV",
                    "9": "Security/Safety"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 2
                },
                "FeatureImage": "new-lower-rates-us-fredericksburg-ha-3211234580-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "New lower rates. Great reviews Elevated deck, Pet friendly, Romantic scenic",
                "PropertySlug": "new-lower-rates-great-reviews-elevated-deck-pet-friendly-romantic-scenic",
                "PropertyType": "Private Vacation Home",
                "RoomSize": 624
            },
            "Partner": {
                "ID": "1234580",
                "OwnerID": "23362313",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Private Vacation Home",
                "URL": "https://www.vrbo.com/search?selected=23362313&regionId=6059096"
            }
        },
        {
            "ID": "BC-8897815",
            "Feed": 11,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "Fredericksburg",
                        "Slug": "usa/texas/fredericksburg",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "fredericksburg"
                        ]
                    },
                    {
                        "Name": "Kerrville",
                        "Slug": "usa/texas/kerrville",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "kerrville"
                        ]
                    }
                ],
                "City": "Kerrville",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "USA",
                "LocationID": "7110",
                "Lat": 30.103796,
                "Lng": -99.115047,
                "Slug": "usa/texas/kerrville"
            },
            "Property": {
                "Amenities": {
                    "1": "Child Friendly",
                    "2": "Internet",
                    "3": "Parking",
                    "4": "Pet Friendly",
                    "5": "Accessibility",
                    "6": "Security/Safety",
                    "7": "Sports/Activities",
                    "8": "Spa",
                    "9": "Fireplace/Heating"
                },
                "Counts": {
                    "Bedroom": 1,
                    "Bathroom": 1,
                    "Occupancy": 4
                },
                "FeatureImage": "cozy-kerrville-guest-cot-us-kerrville-bc-8897815-0.jpg",
                "IsPetFriendly": true,
                "PropertyName": "Cozy Kerrville Guest Cottage Near Guadalupe River!",
                "PropertySlug": "cozy-kerrville-guest-cottage-near-guadalupe-river",
                "PropertyType": "Holiday Homes",
                "RoomSize": 947.22
            },
            "Partner": {
                "ID": "8897815",
                "OwnerID": "",
                "Archived": [],
                "PropertyType": "Holiday Homes",
                "URL": "https://www.booking.com/hotel/us/quaint-kerrville-home-near-guadalupe-river.html?aid=affiliate_id"
            }
        },
        {
            "ID": "HA-321618784",
            "Feed": 12,
            "Published": true,
            "GeoInfo": {
                "Categories": [
                    {
                        "Name": "USA",
                        "Slug": "usa",
                        "Type": "country",
                        "Display": [
                            "usa"
                        ]
                    },
                    {
                        "Name": "Texas",
                        "Slug": "usa/texas",
                        "Type": "state",
                        "Display": [
                            "usa",
                            "texas"
                        ]
                    },
                    {
                        "Name": "New Braunfels",
                        "Slug": "usa/texas/new-braunfels",
                        "Type": "city",
                        "Display": [
                            "usa",
                            "texas",
                            "new-braunfels"
                        ]
                    }
                ],
                "City": "New Braunfels",
                "Country": "USA",
                "CountryCode": "US",
                "Display": "New Braunfels, Texas, United States",
                "LocationID": "6059109",
                "Lat": 29.708882,
                "Lng": -98.110756,
                "Slug": "usa/texas/new-braunfels"
            },
            "Property": {
                "Amenities": {
                    "1": "Air Conditioner",
                    "10": "Pool",
                    "11": "TV",
                    "12": "Wheelchair Accessible",
                    "13": "Accessibility",
                    "14": "Security/Safety",
                    "15": "Sports/Activities",
                    "16": "Wellness Facilities",
                    "17": "Fireplace/Heating",
                    "18": "Entertainment",
                    "19": "Barbecue/Outdoor Cooking",
                    "2": "Balcony/Terrace",
                    "3": "Bedding/Linens",
                    "4": "Child Friendly",
                    "5": "Hot Tub",
                    "6": "Internet",
                    "7": "Kitchen",
                    "8": "Laundry",
                    "9": "Parking"
                },
                "Counts": {
                    "Bedroom": 2,
                    "Bathroom": 2,
                    "Occupancy": 6
                },
                "FeatureImage": "treetop-retreat-on-us-new-braunfels-ha-321618784-0.jpg",
                "IsPetFriendly": false,
                "PropertyName": "Treetop Retreat on the Guadalupe Ultra Private 3rd Floor End Unit",
                "PropertySlug": "treetop-retreat-on-the-guadalupe-ultra-private-3rd-floor-end-unit",
                "PropertyType": "Condo",
                "RoomSize": 940
            },
            "Partner": {
                "ID": "618784",
                "OwnerID": "18994331",
                "Archived": [
                    "VRBO",
                    "EP"
                ],
                "PropertyType": "Condo",
                "URL": "https://www.vrbo.com/search?selected=18994331&regionId=6059109"
            }
        }
    ]
}
```
</details>



#### Error Responses

##### 400 Bad Request

Missing or invalid query parameters.

```json
{
  "error": "location parameter is required"
}
```

```json
{
  "error": "items query parameter must be 'true'"
}
```

##### 401 Unauthorized

Missing or invalid API key.

```json
{
  "error": "x-api-key header is required"
}
```

```json
{
  "error": "invalid API key"
}
```

##### 502 Bad Gateway

External service failure.

```json
{
  "error": "failed to fetch property IDs from location service"
}
```

##### 500 Internal Server Error

Unexpected server errors.

```json
{
  "error": "unexpected server error"
}
```

---
