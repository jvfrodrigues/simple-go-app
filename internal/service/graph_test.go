package service_test

import (
	"reflect"
	"testing"

	"github.com/jvfrodrigues/simple-go-app/internal/model"
	"github.com/jvfrodrigues/simple-go-app/internal/service"
)

func TestNewGraphService(t *testing.T) {
	routes := []model.Route{
		{From: "A", To: "B", Cost: 1},
		{From: "B", To: "C", Cost: 2},
		{From: "A", To: "C", Cost: 4},
	}

	gs := service.NewGraphService(routes)

	if gs == nil {
		t.Fatal("Expected non-nil GraphService")
	}

	if len(gs.Routes) != len(routes) {
		t.Errorf("Expected %d routes, got %d", len(routes), len(gs.Routes))
	}
}

func TestFindBestRoute(t *testing.T) {
	routes := []model.Route{
		{From: "A", To: "B", Cost: 1},
		{From: "B", To: "C", Cost: 2},
		{From: "A", To: "C", Cost: 4},
	}

	gs := service.NewGraphService(routes)

	tests := []struct {
		name          string
		start         string
		end           string
		expectedPath  []string
		expectedCost  int
		expectedFound bool
	}{
		{"Direct route", "A", "B", []string{"A", "B"}, 1, true},
		{"Indirect route", "A", "C", []string{"A", "B", "C"}, 3, true},
		{"No route", "A", "D", nil, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cost := gs.FindBestRoute(tt.start, tt.end)

			if tt.expectedFound {
				if !reflect.DeepEqual(path, tt.expectedPath) {
					t.Errorf("Expected path %v, got %v", tt.expectedPath, path)
				}
				if cost != tt.expectedCost {
					t.Errorf("Expected cost %d, got %d", tt.expectedCost, cost)
				}
			} else {
				if path != nil || cost != 0 {
					t.Errorf("Expected no route, got path %v with cost %d", path, cost)
				}
			}
		})
	}
}

func TestRebuildGraph(t *testing.T) {
	initialRoutes := []model.Route{
		{From: "A", To: "B", Cost: 1},
		{From: "B", To: "C", Cost: 2},
	}

	gs := service.NewGraphService(initialRoutes)

	newRoutes := []model.Route{
		{From: "A", To: "B", Cost: 1},
		{From: "B", To: "C", Cost: 2},
		{From: "A", To: "C", Cost: 4},
	}

	gs.RebuildGraph(newRoutes)

	// Test that the new route is now findable
	path, cost := gs.FindBestRoute("A", "C")
	expectedPath := []string{"A", "B", "C"}
	expectedCost := 3

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("After rebuild, expected path %v, got %v", expectedPath, path)
	}
	if cost != expectedCost {
		t.Errorf("After rebuild, expected cost %d, got %d", expectedCost, cost)
	}
}
