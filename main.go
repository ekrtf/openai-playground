package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	req, err := http.NewRequest(http.MethodGet, "https://api.openai.com/v1/engines", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	client.Do(req)
	// request in a void
	// resp, err := client.Do(req)

	// // body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(body))

	// write list of engines to file
	// err = ioutil.WriteFile("engines.json", body, 0644)

	makeCompletionRequest(apiKey)
}

type BodyRequest struct {
	Model          string `json:"model"`
	ResponseFormat struct {
		Type string `json:"type"`
	} `json:"response_format"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

func makeCompletionRequest(apiKey string) {
	// add JSON body
	reqBody := BodyRequest{
		Model: "gpt-3.5-turbo-0125",
		ResponseFormat: struct {
			Type string `json:"type"`
		}{
			Type: "json_object",
		},
		// pass array of string messages
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: "You are a helpful assistant designed to output JSON.",
			},
		},
	}
	jsonValue, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// set JSON related headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
