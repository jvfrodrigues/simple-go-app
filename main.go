package main

import (
	"fmt"
	"os"

	"github.com/jvfrodrigues/simple-go-app/cmd"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the service you want to access and the input file")
		fmt.Println("Example:")
		fmt.Println("go run main.go cli input.csv")
		fmt.Println("go run main.go api input.csv")
		return
	}

	switch service := os.Args[1]; service {
	case "cli":
		err := cmd.StartCLI()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "api":
		err := cmd.StartAPI()
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
