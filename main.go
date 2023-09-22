package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/sachasmart/terrafir-github-action/types"
)

var (
	apiKey        = flag.String("apiKey", "", "API key Terrafir API.")
	email         = flag.String("email", "", "Email address to send the request to.")
	inputFilePath = flag.String("input", "", "Input file path to the plan that will be assessed.")
	verboseMode   = flag.String("verbose", "", "Verbose mode")
)

func main() {
	flag.Parse()
	preRequestCheck()
	checkEnvironmentVariables(*apiKey, *email, *inputFilePath)
	sendRequest(*apiKey, *email, *inputFilePath)
}

func preRequestCheck() {
	response, err := http.Get(types.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("API is not available")
	}
	color.Green(fmt.Sprintf("API is available %s", response.Status))
}

func sendRequest(apiKey string, email string, inputFilePath string) {
	color.Green(fmt.Sprintf("Using input file: %s", inputFilePath))

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(inputFilePath)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("plan", filepath.Base(inputFilePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/assessment/api", types.URL), payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("email", email)
	req.Header.Add("Authorization", apiKey)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var formattedBody map[string]interface{}
	if err := json.Unmarshal(body, &formattedBody); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(formattedBody)
}

func checkEnvironmentVariables(apiKey string, email string, input string) {
	if apiKey == "" {
		log.Fatalf("API key not provided.")
	}
	if email == "" {
		log.Fatalf("Email address not provided.")
	}
	if input == "" {
		log.Fatalf("Input not provided.")
	}
}
