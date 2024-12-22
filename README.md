# Intelligent Devices Backend

This repository contains a Go-based backend for managing and interacting with intelligent devices. The backend is designed to handle device data, thresholds, and provide RESTful endpoints for communication with the devices.

## Features

- RESTful API for managing thresholds and device data
- Authentication using Basic Auth
- Support for pagination in data retrieval
- JSON responses for seamless integration with devices
- Modular and extensible code structure
- Integration with SQLite database for storing thresholds and device data

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.19 or later
- SQLite database
- A REST client like Postman or cURL for testing API endpoints

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/saumatya/Go-CRUD-API-for-IoT.git
   ```

2. Run the server:
   ```bash
   cd cmd/api
   go run main.go
   ```

## API Endpoints

### Threshold Management

#### Get All Thresholds

**Request:**
```
GET /thresholds?page={page}&rowsPerPage={rowsPerPage}
```

**Example Response:**
```json
[
  {
    "id": 1,
    "sensor_type": "Temperature",
    "min_value": 15.0,
    "max_value": 30.0,
    "updated_at": "2024-12-23T12:00:00Z"
  }
]
```

#### Get Threshold by ID

**Request:**
```
GET /thresholds/{id}
```

**Example Response:**
```json
{
  "id": 1,
  "sensor_type": "Temperature",
  "min_value": 15.0,
  "max_value": 30.0,
  "updated_at": "2024-12-23T12:00:00Z"
}
```

#### Create a New Threshold

**Request:**
```
POST /thresholds
```

**Example Payload:**
```json
{
  "sensor_type": "Humidity",
  "min_value": 20.0,
  "max_value": 60.0
}
```

#### Update a Threshold

**Request:**
```
PUT /thresholds/{id}
```

**Example Payload:**
```json
{
  "min_value": 18.0,
  "max_value": 28.0
}
```

#### Delete a Threshold

**Request:**
```
DELETE /thresholds/{id}
```

**Example Response:**
```json
{
  "message": "Threshold successfully deleted"
}
```