package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Mock implementation of DataService for testing purposes, always returns a successful response and Data object(s) *
type MockDataServiceSuccessful struct{}

func (m *MockDataServiceSuccessful) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return []*models.Data{
		{
			ID:          1,
			DeviceID:    "device1",
			DeviceName:  "device1",
			TemperatureValue: 0.0,
			HumidityValue: 0.0,
			Type:        "type1",
			DateTime:    "2021-01-01 00:00:00",
		},
		{
			ID:          2,
			DeviceID:    "device2",
			DeviceName:  "device2",
			TemperatureValue: 0.0,
			HumidityValue: 0.0,
			Type:        "type2",
			DateTime:    "2021-01-01 00:00:00",
		},
	}, nil
}

func (m *MockDataServiceSuccessful) ReadOne(id int, ctx context.Context) (*models.Data, error) {
	return &models.Data{
		ID:          1,
		DeviceID:    "device1",
		DeviceName:  "device1",
		TemperatureValue: 0.0,
		HumidityValue: 0.0,
		Type:        "type1",
		DateTime:    "2021-01-01 00:00:00",

	}, nil
}

func (m *MockDataServiceSuccessful) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) Delete(data *models.Data, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) Create(data *models.Data, ctx context.Context) error {
	return nil
}

// * Threshold-specific mocks *
func (m *MockDataServiceSuccessful) CreateThreshold(threshold *models.Threshold, ctx context.Context) error {
	// Return nil to signify successful creation
	return nil
}

func (m *MockDataServiceSuccessful) ReadThreshold(id int, ctx context.Context) (*models.Threshold, error) {
	// Return a sample threshold object
	return &models.Threshold{
		ID:      id,
		MinValue: 10.0,
		MaxValue: 50.0,
		SensorType:    "Temperature",
	}, nil
}
func (m *MockDataServiceSuccessful) UpdateThreshold(threshold *models.Threshold, ctx context.Context) (int64, error) {
	// Return 1 to signify a successful update with a placeholder record ID
	return 1, nil
}

func (m *MockDataServiceSuccessful) DeleteThreshold(id int, ctx context.Context) (int64, error) {
	// Return 1 to signify a successful delete of the threshold
	return 1, nil
}

func (m *MockDataServiceSuccessful) GetAllThresholds(page, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
	// Return a list of sample thresholds
	return []*models.Threshold{
		{ID: 1, MinValue: 10.0, MaxValue: 50.0, SensorType: "Temperature"},
		{ID: 2, MinValue: 20.0, MaxValue: 60.0, SensorType: "Humidity"},
	}, nil
}

func (m *MockDataServiceSuccessful) ValidateData(data *models.Data) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns empty data *

type MockDataServiceNotFound struct{}

func (m *MockDataServiceNotFound) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return []*models.Data{}, nil
}

func (m *MockDataServiceNotFound) ReadOne(id int, ctx context.Context) (*models.Data, error) {
	return nil, nil
}

func (m *MockDataServiceNotFound) Create(data *models.Data, ctx context.Context) error {
	return nil
}

func (m *MockDataServiceNotFound) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) Delete(data *models.Data, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) CreateThreshold(threshold *models.Threshold, ctx context.Context) error {
	// No action, implicitly returning nil for error
	return nil
}

func (m *MockDataServiceNotFound) ReadThreshold(id int, ctx context.Context) (*models.Threshold, error) {
	// Simulate a scenario where the threshold is not found
	return nil, nil
}

func (m *MockDataServiceNotFound) UpdateThreshold(threshold *models.Threshold, ctx context.Context) (int64, error) {
	// Simulate no records affected for an update attempt
	return 0, nil
}

func (m *MockDataServiceNotFound) DeleteThreshold(id int, ctx context.Context) (int64, error) {
	// Simulate no records affected for delete attempt
	return 0, nil
}

func (m *MockDataServiceNotFound) GetAllThresholds(page, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
	// Return empty list for no thresholds found
	return []*models.Threshold{}, nil
}

func (m *MockDataServiceNotFound) ValidateData(data *models.Data) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns an error *
type MockDataServiceError struct{}

func (m *MockDataServiceError) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) ReadOne(id int, ctx context.Context) (*models.Data, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) Create(data *models.Data, ctx context.Context) error {
	return DataError{Message: "Error creating data."}
}

func (m *MockDataServiceError) Update(data *models.Data, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error updating data."}
}

func (m *MockDataServiceError) Delete(data *models.Data, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error deleting data."}
}
// Mock for CreateThreshold - returning a DataError
func (m *MockDataServiceError) CreateThreshold(threshold *models.Threshold, ctx context.Context) error {
	return DataError{Message: "Error creating threshold."}
}

// Mock for ReadThreshold - returning a DataError
func (m *MockDataServiceError) ReadThreshold(id int, ctx context.Context) (*models.Threshold, error) {
	return nil, DataError{Message: "Error reading threshold."}
}

// Mock for UpdateThreshold - returning a DataError
func (m *MockDataServiceError) UpdateThreshold(threshold *models.Threshold, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error updating threshold."}
}

// Mock for DeleteThreshold - returning a DataError
func (m *MockDataServiceError) DeleteThreshold(id int, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error deleting threshold."}
}

// Mock for GetAllThresholds - returning a DataError
func (m *MockDataServiceError) GetAllThresholds(page, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
	return nil, DataError{Message: "Error retrieving thresholds."}
}



func (m *MockDataServiceError) ValidateData(data *models.Data) error {
	return nil
}
