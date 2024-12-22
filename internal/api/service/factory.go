package service

import (
	"context"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/DAL/SQLite"
	service "goapi/internal/api/service/data"
	"log"
)

type DataServiceType int

const (
	SQLiteDataService DataServiceType = iota
)

type ServiceFactory struct {
	db     DAL.SQLDatabase
	logger *log.Logger
	ctx    context.Context
}

// * Factory for creating data service *
func NewServiceFactory(db DAL.SQLDatabase, logger *log.Logger, ctx context.Context) *ServiceFactory {
	return &ServiceFactory{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}

func (sf *ServiceFactory) CreateDataService(serviceType DataServiceType) (*service.DataServiceSQLite, error) {
	switch serviceType {
	case SQLiteDataService:
		// Create both repositories
		dataRepo, err := SQLite.NewDataRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		thresholdRepo, err := SQLite.NewThresholdRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		// Create the DataServiceSQLite with both repositories
		ds := service.NewDataServiceSQLite(dataRepo, thresholdRepo)
		return ds, nil
	default:
		return nil, service.DataError{Message: "Invalid data service type."}
	}
}
