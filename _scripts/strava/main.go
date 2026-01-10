package main

import (
	"log"
	"os"
)

func main() {
	// Read Strava API credentials from environment variables
	clientId := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	refreshToken := os.Getenv("STRAVA_REFRESH_TOKEN")

	if clientId == "" || clientSecret == "" || refreshToken == "" {
		log.Fatal("Missing required environment variables: STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, STRAVA_REFRESH_TOKEN")
	}

	// Refresh the access token
	tokenResponse, err := RefreshAccessToken(clientId, clientSecret, refreshToken)
	if err != nil {
		log.Fatal(err)
	}

	// Get activities using the fresh access token
	activities, err := GetActivities(tokenResponse.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	// Write activities to Jekyll collection
	siteBaseDir := "../../_activities"
	err = WriteActivities(activities, siteBaseDir)
	if err != nil {
		log.Fatal(err)
	}
}
