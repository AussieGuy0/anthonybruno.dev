package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var stravaApiUrl = "https://www.strava.com/api/v3"
var stravaOAuthUrl = "https://www.strava.com/oauth/token"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
}

type Activity struct {
	Id                 int64   `json:"id"`
	Name               string  `json:"name"`
	Distance           float64 `json:"distance"`             // in meters
	MovingTime         int     `json:"moving_time"`          // in seconds
	ElapsedTime        int     `json:"elapsed_time"`         // in seconds
	TotalElevationGain float64 `json:"total_elevation_gain"` // in meters
	Type               string  `json:"type"`                 // e.g., "Run", "Ride"
	StartDate          string  `json:"start_date"`           // ISO 8601 formatted date time
	StartDateLocal     string  `json:"start_date_local"`
	AverageSpeed       float64 `json:"average_speed"` // in meters per second
	MaxSpeed           float64 `json:"max_speed"`     // in meters per second
}

func RefreshAccessToken(clientId, clientSecret, refreshToken string) (*TokenResponse, error) {
	log.Println("Refreshing Strava access token")

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("refresh_token", refreshToken)
	data.Set("grant_type", "refresh_token")

	res, err := http.PostForm(stravaOAuthUrl, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New("Failed to refresh token. Status: " + strconv.Itoa(res.StatusCode) + ". Body: " + string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully refreshed access token (expires at: %d)", tokenResponse.ExpiresAt)
	return &tokenResponse, nil
}

func GetActivities(accessToken string) ([]Activity, error) {
	res, err := makeActivitiesRequest(accessToken)
	if err != nil {
		return nil, err
	}
	activities, err := parseActivities(res)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (activity Activity) StartDateTime() (time.Time, error) {
	// Parse the ISO 8601 formatted date
	return time.Parse(time.RFC3339, activity.StartDate)
}

func (activity Activity) DistanceKm() float64 {
	return activity.Distance / 1000.0
}

func (activity Activity) MovingTimeFormatted() string {
	duration := time.Duration(activity.MovingTime) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	return fmt.Sprintf("%dm %ds", minutes, seconds)
}

func (activity Activity) AverageSpeedKmh() float64 {
	return activity.AverageSpeed * 3.6
}

func (activity Activity) PaceMinPerKm() string {
	if activity.AverageSpeed == 0 {
		return "0:00"
	}
	secondsPerKm := 1000.0 / activity.AverageSpeed
	minutes := int(secondsPerKm / 60)
	seconds := int(secondsPerKm) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func makeActivitiesRequest(accessToken string) (*http.Response, error) {
	url := stravaApiUrl + "/athlete/activities"

	log.Println("Getting activities from Strava")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	// Add query parameters to get more activities (default is 30)
	q := req.URL.Query()
	q.Add("per_page", "200") // Get up to 200 activities
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New("Non 200 status code: " + strconv.Itoa(res.StatusCode) + ". Response: " + string(body))
	}
	return res, nil
}

func parseActivities(res *http.Response) ([]Activity, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var activities []Activity
	err = json.Unmarshal(body, &activities)
	if err != nil {
		return nil, err
	}
	return activities, nil
}
