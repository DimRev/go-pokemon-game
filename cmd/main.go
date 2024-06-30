package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DimRev/go-pokemon-game/internal/game"
)

func main() {
	help := false
	start := false

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") {
			switch arg {
			case "--help":
				help = true
			case "--start":
				start = true
			}
		} else if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-h":
				help = true
			case "-s":
				start = true
			default:
				fmt.Printf("flag provided but not defined: %s\n", arg)
				printUsage()
				os.Exit(2)
			}
		}
	}

	flag.Usage = printUsage

	if help {
		flag.Usage()
	}

	if start {
		game.Test()
	}
}

func printUsage() {
	fmt.Println("Usage of the CLI game:")
	fmt.Println("  --help   -h :  Display help")
	fmt.Println("  --start  -s :  Start Game")
}

func getLocations() error {
	url := "https://pokeapi.co/api/v2/location/"

	type Response struct {
		Count    int         `json:"count"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal JSON response into struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	// Print the structured response
	fmt.Printf("Count: %d\n", response.Count)
	fmt.Printf("Next: %s\n", response.Next)
	fmt.Printf("Previous: %v\n", response.Previous)
	fmt.Println("Locations:")
	for _, location := range response.Results {
		fmt.Printf(" - %s: %s\n", location.Name, location.URL)
	}

	return nil
}
