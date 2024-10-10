package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	//   "github.com/dghubble/oauth2"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the credentials from environment variables
	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecretKey := os.Getenv("TWITTER_API_SECRET_KEY")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	bearerToken := os.Getenv("TWITTER_BEARER_TOKEN")

	//Check if all variables are set
	if apiKey == "" || apiSecretKey == "" || accessToken == "" || accessTokenSecret == "" || bearerToken == "" {
		log.Fatal("One or more environment variables are not set.")
	}

	// Create the OAuth1 configuration
	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create an HTTP client with OAuth1
	httpClient := config.Client(oauth1.NoContext, token)

	// Define the tweet payload
	tweetText := "Hello Twitter! This is my third tweet from the Twitter API v2."
	payload := map[string]interface{}{
		"text": tweetText,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling payload: %v", err)
	}

	// Create the request to post a tweet
	url := "https://api.twitter.com/2/tweets"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the required headers
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Failed to post tweet: %s", resp.Status)
	}

	// Read and print the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	fmt.Printf("Posted Tweet ID: %v\n", result["data"].(map[string]interface{})["id"])

}
