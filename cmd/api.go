package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jvfrodrigues/simple-go-app/internal/infra"
	"github.com/jvfrodrigues/simple-go-app/internal/service"
)

func StartAPI() error {
	inputFile := os.Args[2]
	csvInterpreter := &infra.CSVInterpreter{
		Filename: inputFile,
	}
	routes, err := csvInterpreter.ReadRoutes()
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	graphService := service.NewGraphService(routes)
	apiService := service.NewAPIService(graphService, csvInterpreter, inputFile)

	http.HandleFunc("/route", apiService.HandleRouteQuery)
	http.HandleFunc("/route/add", apiService.HandleAddRoute)

	fmt.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}
