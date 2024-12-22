package data

import (
	"context"
	"goapi/internal/api/repository/models"
	"time"
)

// * Implementation of DataService for SQLite database *
type DataServiceSQLite struct {
	repo models.DataRepository
	thresholdRepo    models.ThresholdRepository 
}

func NewDataServiceSQLite(repo models.DataRepository, thresholdRepo models.ThresholdRepository) *DataServiceSQLite {
	return &DataServiceSQLite{
		repo: repo,
		thresholdRepo: thresholdRepo,
	}
}

func (ds *DataServiceSQLite) Create(data *models.Data, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return DataError{Message: "InvalMockDataServiceSuccessfulid data."}
	}
	return ds.repo.Create(data, ctx)
}

func (ds *DataServiceSQLite) ReadOne(id int, ctx context.Context) (*models.Data, error) {

	data, err := ds.repo.ReadOne(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	// Tehdään datalle jotain, päätellään datasta jotain!!!
	// Tämä ohjaa toimintaa älykkäästi, esim. jos data on tietynlaista, niin tehdään jotain

	return data, nil
}

func (ds *DataServiceSQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return ds.repo.ReadMany(page, rowsPerPage, ctx)
}

func (ds *DataServiceSQLite) Update(data *models.Data, ctx context.Context) (int64, error) {

	if err := ds.ValidateData(data); err != nil {
		return 0, DataError{Message: "Invalid data: " + err.Error()}
	}
	return ds.repo.Update(data, ctx)
}

func (ds *DataServiceSQLite) Delete(data *models.Data, ctx context.Context) (int64, error) {
	return ds.repo.Delete(data, ctx)
}
func (ds *DataServiceSQLite) DeleteThreshold(id int, ctx context.Context) (int64, error) {
    // Create a threshold object with just the ID
    threshold := &models.Threshold{ID: id}
    
    // Call the repository's Delete method
    result, err := ds.thresholdRepo.Delete(threshold, ctx)
    if err != nil {
        return 0, err
    }
    return result, nil
}


func (ds *DataServiceSQLite) ValidateData(data *models.Data) error {
	var errMsg string
	if data.DeviceID == "" || len(data.DeviceID) > 50 {
		errMsg += "DeviceID is required and must be less than 50 characters. "
	}
	if len(data.DeviceName) > 50 {
		errMsg += "DeviceName must be less than 50 characters. "
	}
	if len(data.Type) > 20 {
		errMsg += "Type must be less than 20 characters. "
	}
	if data.TemperatureValue > 100 {
		errMsg += "Temperature must be less than 100 degrees "
	}
	if data.HumidityValue > 100 {
		errMsg += "Humidity must be less than 100 %. "
	}
	_, err := time.Parse("2006-01-02T15:04:05Z", data.DateTime)
	if err != nil {
		errMsg += "DateTime must be in the format: 2021-01-01T12:00:00Z. "
	}
	if errMsg != "" {
		return DataError{Message: errMsg}
	}
	return nil
}
func (ds *DataServiceSQLite) CreateThreshold(threshold *models.Threshold, ctx context.Context) error {
	// Example logic for creating a Threshold
	if threshold.SensorType == "" {
		return DataError{Message: "SensorType is required."}
	}
	if threshold.MinValue >= threshold.MaxValue {
		return DataError{Message: "MinValue should be less than MaxValue."}
	}
	
	return ds.thresholdRepo.Create(threshold, ctx)
}
func (ds *DataServiceSQLite) GetAllThresholds(page, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
    // Call the repository method to get all thresholds with pagination
    thresholds, err := ds.thresholdRepo.ReadMany(page, rowsPerPage, ctx)
    if err != nil {
        return nil, err
    }
    return thresholds, nil
}

// Read a single threshold by ID
func (ds *DataServiceSQLite) ReadThreshold(id int, ctx context.Context) (*models.Threshold, error) {
    threshold, err := ds.thresholdRepo.ReadOne(id, ctx)
    if err != nil {
        return nil, err
    }
    return threshold, nil
}

// Get all thresholds with pagination
// func (ds *DataServiceSQLite) GetAllThresholds(page, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
//     thresholds, err := ds.thresholdRepo.ReadMany(page, rowsPerPage, ctx)
//     if err != nil {
//         return nil, err
//     }
//     return thresholds, nil
// }

// Create a new threshold
// func (ds *DataServiceSQLite) CreateThreshold(threshold *models.Threshold, ctx context.Context) error {
//     if err := ds.validateThreshold(threshold); err != nil {
//         return err
//     }
//     return ds.thresholdRepo.Create(threshold, ctx)
// }

// Update an existing threshold
func (ds *DataServiceSQLite) UpdateThreshold(threshold *models.Threshold, ctx context.Context) (int64, error) {
    if err := ds.validateThreshold(threshold); err != nil {
        return 0, err
    }
    return ds.thresholdRepo.Update(threshold, ctx)
}

// Helper function to validate threshold data
func (ds *DataServiceSQLite) validateThreshold(threshold *models.Threshold) error {
    if threshold.SensorType == "" {
        return DataError{Message: "SensorType is required"}
    }
    if threshold.MinValue >= threshold.MaxValue {
        return DataError{Message: "MinValue must be less than MaxValue"}
    }
    return nil
}