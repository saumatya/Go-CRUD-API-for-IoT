package models

import "context"

type Threshold struct {
    ID         int     `json:"id"`
    SensorType string  `json:"sensor_type"`
    MinValue   float64 `json:"min_value"`
    MaxValue   float64 `json:"max_value"`
    UpdatedAt  string  `json:"updated_at"`
}

type ThresholdRepository interface {
    Create(threshold *Threshold, ctx context.Context) error
    ReadOne(id int, ctx context.Context) (*Threshold, error)
    ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Threshold, error)
    Update(threshold *Threshold, ctx context.Context) (int64, error)
    Delete(threshold *Threshold, ctx context.Context) (int64, error)
}
