package main

import (
	"bytes"
	"encoding/json"
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
	apiKey      = getEnv("INPUT_APIKEY", "")
	email       = getEnv("INPUT_EMAIL", "")
	path        = getEnv("INPUT_PATH", "")
	verboseMode = getEnv("verboseMode", "")
)

func getEnv(key, defaultValue string) string {
	fmt.Println(os.Environ())
	value := os.Getenv(key)
	if value == "" {
		fmt.Print("Using default value for ", key, ": ", defaultValue, "\n")
		return defaultValue
	}
	return value
}+

func main() {
	preRequestCheck()
	sendRequest(apiKey, email)
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

func sendRequest(apiKey string, email string) {
	color.Green(fmt.Sprintf("Using input file: %s", "./input.json"))

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("./input.json")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("plan", filepath.Base("./input.json"))
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
	formattedBody(body)

}

func formattedBody(body []byte) {
	var formattedBody map[string]interface{}
	if err := json.Unmarshal(body, &formattedBody); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(formattedBody)
}
