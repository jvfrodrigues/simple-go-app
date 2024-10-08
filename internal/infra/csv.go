package infra

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jvfrodrigues/simple-go-app/internal/model"
)

type CSVInterpreter struct {
	Filename string
}

func (ci *CSVInterpreter) ReadRoutes() ([]model.Route, error) {
	file, err := os.Open(ci.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var routes []model.Route

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		cost, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("invalid cost in CSV: %v", err)
		}

		routes = append(routes, model.Route{
			From: record[0],
			To:   record[1],
			Cost: cost,
		})
	}

	return routes, nil
}

func (ci *CSVInterpreter) WriteNewRoute(filename string, route model.Route) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{route.From, route.To, strconv.Itoa(route.Cost)})
}
