package models

import "context"

type Data struct {
	ID          int     `json:"id"`
	DeviceID    string  `json:"device_id"`
	DeviceName  string  `json:"device_name"`
	TemperatureValue       float64 `json:"temp_value"`
	HumidityValue       float64 `json:"humi_value"`
	Type        string  `json:"type"`
	DateTime    string  `json:"date_time"`
}

type DataRepository interface {
	Create(Data *Data, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*Data, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	Update(data *Data, ctx context.Context) (int64, error)
	Delete(data *Data, ctx context.Context) (int64, error)
}
