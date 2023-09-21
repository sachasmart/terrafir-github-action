package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/sachasmart/terrafir-github-action/types"
)

var (
	apiKey      = flag.String("api-key", os.Getenv("API_KEY"), "API key Terrafir API.")
	email       = flag.String("email", os.Getenv("EMAIL"), "Email address to send the request to.")
	input       = flag.String("input", os.Getenv("INPUT"), "Input to send to the API.")
	verboseMode = flag.String("verbose", os.Getenv("VERBOSE"), "Verbose mode")
)

func main() {
	godotenv.Load(".env")
	flag.Parse()
	checkEnvironmentVariables(*apiKey, *email, *input)
	sendRequest(*apiKey, *email, *input)
}

func preRequestCheck() {
	response, err := http.Get(types.URL)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatal("API is not available")
	}
	fmt.Println("API is available")
}

func sendRequest(apiKey string, email string, input string) {
	preRequestCheck()
	response, err := http.Post(types.URL, "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Body)
}

func checkEnvironmentVariables(apiKey string, email string, input string) {
	if apiKey == "" {
		log.Fatalf("Environment variable 'API_KEY' not found.")
	}
	_, err := url.Parse(apiKey)
	if err != nil {
		log.Fatalf("Environment variable 'API_KEY' is not a valid.")
	}
}
