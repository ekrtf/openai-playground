package main

import (
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

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
