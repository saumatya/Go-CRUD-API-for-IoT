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

func TestPostInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("POST", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	data.PostHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostErrorCreatingData(t *testing.T) {

	req, err := http.NewRequest("POST", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Data{
		ID:          1,
		DeviceID:    "device1",
		DeviceName:  "device1",
		TemperatureValue:       10.0,
		HumidityValue: 10.0,
		Type:        "type1",
		DateTime:    "2021-01-01 00:00:00",

	})

	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))
	rr := httptest.NewRecorder()

	data.PostHandler(rr, req, log.Default(), &service.MockDataServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Error creating data."}` // * This message is passed from the MockDataServiceError
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestPostSuccessful(t *testing.T) {
	req, err := http.NewRequest("POST", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}
	dataJSON, _ := json.Marshal(models.Data{
		ID:               1,
		DeviceID:         "device1",
		DeviceName:       "device1",
		TemperatureValue: 10.0,  
		HumidityValue:    10.0,  
		Type:             "type1",
		DateTime:         "2021-01-01 00:00:00",
	})
	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))
	rr := httptest.NewRecorder()
	data.PostHandler(rr, req, log.Default(), &service.MockDataServiceSuccessful{})
	// Check the status code is HTTP 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	expected := `{"id":1,"device_id":"device1","device_name":"device1","temp_value":10,"humi_value":10,"type":"type1","date_time":"2021-01-01 00:00:00"}`

	// Check if the response body matches the expected output
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}