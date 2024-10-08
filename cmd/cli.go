package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/jvfrodrigues/simple-go-app/internal/infra"
	"github.com/jvfrodrigues/simple-go-app/internal/service"
)

func StartCLI() error {
	inputFile := os.Args[2]
	csvInterpreter := &infra.CSVInterpreter{
		Filename: inputFile,
	}
	routes, err := csvInterpreter.ReadRoutes()
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	graphService := service.NewGraphService(routes)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please enter the route: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		fromTo := strings.Split(input, "-")
		if len(fromTo) != 2 {
			fmt.Println("Invalid input. Please use format: FROM-TO")
			continue
		}

		from, to := fromTo[0], fromTo[1]
		bestRoute, totalCost := graphService.FindBestRoute(from, to)

		if len(bestRoute) == 0 {
			fmt.Println("No route found")
		} else {
			fmt.Printf("best route: %s > $%d\n", strings.Join(bestRoute, " - "), totalCost)
		}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	return nil
}
