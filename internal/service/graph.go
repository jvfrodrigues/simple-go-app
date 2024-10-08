package service

import (
	"math"

	"github.com/jvfrodrigues/simple-go-app/internal/model"
)

type GraphService struct {
	Routes  []model.Route
	graphed model.GraphedRoutes
}

func NewGraphService(routes []model.Route) *GraphService {
	graph := buildGraph(routes)
	return &GraphService{
		Routes:  routes,
		graphed: graph,
	}
}

func (gc *GraphService) RebuildGraph(updatedRoutes []model.Route) {
	gc.graphed = buildGraph(updatedRoutes)
}

func (gs *GraphService) FindBestRoute(start, end string) ([]string, int) {
	_, startOk := gs.graphed[start]
	_, endOk := gs.graphed[end]

	if !startOk || !endOk {
		return nil, 0
	}

	distances := make(map[string]int)
	previous := make(map[string]string)
	unvisited := make(map[string]bool)

	for node := range gs.graphed {
		distances[node] = math.MaxInt32
		previous[node] = ""
		unvisited[node] = true
	}
	distances[start] = 0

	for len(unvisited) > 0 {
		var current string
		minDist := math.MaxInt32

		for node := range unvisited {
			if distances[node] < minDist {
				minDist = distances[node]
				current = node
			}
		}

		if current == end {
			break
		}

		delete(unvisited, current)

		for neighbor, cost := range gs.graphed[current] {
			alt := distances[current] + cost
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				previous[neighbor] = current
			}
		}
	}

	if distances[end] == math.MaxInt32 {
		return nil, 0
	}

	path := []string{}
	for node := end; node != ""; node = previous[node] {
		path = append([]string{node}, path...)
	}

	return path, distances[end]
}

func buildGraph(routes []model.Route) model.GraphedRoutes {
	graph := make(model.GraphedRoutes)

	for _, route := range routes {
		if _, exists := graph[route.From]; !exists {
			graph[route.From] = make(map[string]int)
		}
		graph[route.From][route.To] = route.Cost

		// Assuming routes are bidirectional
		if _, exists := graph[route.To]; !exists {
			graph[route.To] = make(map[string]int)
		}
		graph[route.To][route.From] = route.Cost
	}

	return graph
}
