package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/plivo/plivo-go"
)

func main() {
	var err error
	// Load .env file
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Set environment variables
	auth_id := os.Getenv("PLIVO_AUTH_ID")
	auth_token := os.Getenv("PLIVO_AUTH_TOKEN")
	from_num := os.Getenv("PLIVO_FROM")
	to_num := os.Getenv("PLIVO_TO")

	// Create http client
	client, err := plivo.NewClient(auth_id, auth_token, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Messages.Create(
		plivo.MessageCreateParams{
			Src: from_num,
			Dst: to_num,
			Text: "Plivo Go!",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}