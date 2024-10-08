package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jvfrodrigues/simple-go-app/internal/infra"
	"github.com/jvfrodrigues/simple-go-app/internal/model"
)

type APIService struct {
	csvInterpreter *infra.CSVInterpreter
	graphService   *GraphService
	inputFile      string
}

func NewAPIService(graphService *GraphService, csvInterpreter *infra.CSVInterpreter, inputFile string) *APIService {
	return &APIService{
		graphService:   graphService,
		csvInterpreter: csvInterpreter,
		inputFile:      inputFile,
	}
}

func (s *APIService) HandleRouteQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" {
		http.Error(w, "Missing 'from' or 'to' parameter", http.StatusBadRequest)
		return
	}

	bestRoute, totalCost := s.graphService.FindBestRoute(from, to)

	response := map[string]interface{}{
		"route": bestRoute,
		"cost":  totalCost,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *APIService) HandleAddRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newRoute model.Route
	err = json.Unmarshal(body, &newRoute)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if newRoute.From == "" || newRoute.To == "" || newRoute.Cost <= 0 {
		http.Error(w, "Invalid route data", http.StatusBadRequest)
		return
	}

	err = s.csvInterpreter.WriteNewRoute(s.inputFile, newRoute)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding route: %v", err), http.StatusInternalServerError)
		return
	}

	updated, err := s.csvInterpreter.ReadRoutes()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading routes: %v", err), http.StatusInternalServerError)
		return
	}

	s.graphService.RebuildGraph(updated)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Route added successfully")
}
