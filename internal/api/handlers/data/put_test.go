package data_test

import (
	"encoding/json"
	"goapi/internal/api/handlers/data"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/data"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPutInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("PUT", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	data.PutHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPutHandlerError(t *testing.T) {

	req, err := http.NewRequest("PUT", "/data", strings.NewReader(`{"id": 1, "device_id": "device_id", "device_name": "device_name", "temp_value": 1.0, "humi_value": 1.0, "type": "type", "date_time": "2020-01-01T00:00:00Z"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	data.PutHandler(rr, req, log.Default(), &service.MockDataServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Error updating data."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPutDataNotFound(t *testing.T) {

	req, err := http.NewRequest("PUT", "/data", strings.NewReader(`{
		"id": 1,
		"device_id": "device_id",
		"device_name": "device_name",
		"temp_value": 10.0,
		"humi_value": 10.0,
		"type": "type",
		"date_time": "2020-01-01T00:00:00Z"
	}`))
		if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	data.PutHandler(rr, req, log.Default(), &service.MockDataServiceNotFound{})

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	expected := `{"error": "Resource not found."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPutHandlerSuccess(t *testing.T) {

	req, err := http.NewRequest("PUT", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Data{
		ID:               1,
		DeviceID:         "device_id",
		DeviceName:       "device_name",
		TemperatureValue: 10.0,  // Keep as float in input
		HumidityValue:    10.0,  // Keep as float in input
		Type:             "type",
		DateTime:         "2020-01-01T00:00:00Z",
	})

	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))

	rr := httptest.NewRecorder()

	// Call the handler
	data.PutHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})

	// Check for the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Modify the expected to use integers for temp_value and humi_value
	expected := `{"id":1,"device_id":"device_id","device_name":"device_name","temp_value":10,"humi_value":10,"type":"type","date_time":"2020-01-01T00:00:00Z"}`

	// Check if the response body matches the expected output
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
