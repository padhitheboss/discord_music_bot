package config

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var YoutubeService *youtube.Service

func connectToYoutube() (*youtube.Service, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to read .env file")
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	ctx := context.Background()

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return &youtube.Service{}, err
	}
	return service, nil
}
func init() {
	var err error
	YoutubeService, err = connectToYoutube()
	if err != nil {
		log.Panic(err)
		return
	}
	log.Println("Connected to Youtube Server")
}
