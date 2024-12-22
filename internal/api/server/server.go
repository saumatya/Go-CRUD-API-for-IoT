package server

import (
	"context"
	"goapi/internal/api/handlers/data"
	"goapi/internal/api/middleware"
	"goapi/internal/api/service"
	"log"
	"net/http"
)

type Server struct {
	ctx        context.Context
	HTTPServer *http.Server
	logger     *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {
	mux := http.NewServeMux()

	// Setup data-related handlers
	err := setupDataHandlers(mux, sf, logger)
	if err != nil {
		logger.Fatalf("Error setting up data handlers: %v", err)
	}

	// Setup threshold-related handlers
	err = setupThresholdHandlers(mux, sf, logger)
	if err != nil {
		logger.Fatalf("Error setting up threshold handlers: %v", err)
	}

	middlewares := []middleware.Middleware{
		middleware.BasicAuthenticationMiddleware,
		middleware.CommonMiddleware,
	}

	return &Server{
		ctx:    ctx,
		logger: logger,
		HTTPServer: &http.Server{
			Handler: middleware.ChainMiddleware(mux, middlewares...),
		},
	}
}

func (api *Server) Shutdown() error {
	api.logger.Println("Gracefully shutting down server...")
	return api.HTTPServer.Shutdown(api.ctx)
}

func (api *Server) ListenAndServe(addr string) error {
	api.HTTPServer.Addr = addr
	return api.HTTPServer.ListenAndServe()
}

// * REST API handlers for Data *
func setupDataHandlers(mux *http.ServeMux, sf *service.ServiceFactory, logger *log.Logger) error {
	ds, err := sf.CreateDataService(service.SQLiteDataService)
	if err != nil {
		return err
	}

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Handle the OPTIONS request to allow for CORS or pre-flight checks
		if r.Method == "OPTIONS" {
			data.OptionsHandler(w, r)
		} else if r.Method == "POST" {
			data.PostHandler(w, r, logger, ds)
		} else if r.Method == "PUT" {
			data.PutHandler(w, r, logger, ds)
		} else if r.Method == "GET" {
			data.GetHandler(w, r, logger, ds)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Use a separate route for handling the ID-based actions
	mux.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data.GetByIDHandler(w, r, logger, ds)
		} else if r.Method == "DELETE" {
			data.DeleteHandler(w, r, logger, ds)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return nil
}

// * REST API handlers for Threshold *
func setupThresholdHandlers(mux *http.ServeMux, sf *service.ServiceFactory, logger *log.Logger) error {
	ds, err := sf.CreateDataService(service.SQLiteDataService)
	if err != nil {
		return err
	}

	mux.HandleFunc("/threshold", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data.PostThresholdHandler(w, r, logger, ds)
		} else if r.Method == "GET" {
			data.GetThresholdHandler(w, r, logger, ds)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})


	mux.HandleFunc("/threshold/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			data.DeleteThresholdHandler(w, r, logger, ds)
		} else if r.Method == "GET" {
			data.GetThresholdHandler(w, r, logger, ds)
		}else if r.Method == "PUT" {
			data.UpdateThresholdHandler(w, r, logger, ds) // Only handle PUT with ID
		}else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return nil
}
