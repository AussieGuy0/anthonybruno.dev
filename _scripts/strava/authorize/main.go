package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Load client credentials from environment or prompt
	clientId := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")

	if clientId == "" {
		fmt.Print("Enter your Strava Client ID: ")
		clientId, _ = reader.ReadString('\n')
		clientId = strings.TrimSpace(clientId)
	}

	if clientSecret == "" {
		fmt.Print("Enter your Strava Client Secret: ")
		clientSecret, _ = reader.ReadString('\n')
		clientSecret = strings.TrimSpace(clientSecret)
	}

	// Generate authorization URL
	authUrl := fmt.Sprintf(
		"https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost&approval_prompt=force&scope=activity:read_all",
		clientId,
	)

	fmt.Println("\n=== Strava Authorization ===")
	fmt.Println("\nStep 1: Open this URL in your browser:")
	fmt.Println(authUrl)
	fmt.Println("\nStep 2: Authorize the application")
	fmt.Println("\nStep 3: You'll be redirected to a URL like:")
	fmt.Println("http://localhost/?state=&code=XXXXX&scope=read,activity:read_all")
	fmt.Print("\nPaste the full redirect URL here: ")

	redirectUrl, _ := reader.ReadString('\n')
	redirectUrl = strings.TrimSpace(redirectUrl)

	// Parse the code from the redirect URL
	parsedUrl, err := url.Parse(redirectUrl)
	if err != nil {
		log.Fatal("Failed to parse redirect URL: ", err)
	}

	code := parsedUrl.Query().Get("code")
	if code == "" {
		log.Fatal("No authorization code found in the URL")
	}

	fmt.Println("\n=== Exchanging authorization code for tokens ===")

	// Exchange code for tokens
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")

	res, err := http.PostForm("https://www.strava.com/oauth/token", data)
	if err != nil {
		log.Fatal("Failed to exchange code: ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed to read response: ", err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Failed to get tokens. Status: %d. Body: %s", res.StatusCode, string(body))
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatal("Failed to parse response: ", err)
	}

	fmt.Println("\n=== Success! ===")
	fmt.Println("\nYour tokens:")
	fmt.Printf("Access Token: %s\n", tokenResponse["access_token"])
	fmt.Printf("Refresh Token: %s\n", tokenResponse["refresh_token"])
	fmt.Printf("Expires At: %v\n", tokenResponse["expires_at"])

	fmt.Println("\n=== Update your .env file ===")
	fmt.Println("\nAdd these lines to _scripts/strava/.env:")
	fmt.Printf("STRAVA_CLIENT_ID=%s\n", clientId)
	fmt.Printf("STRAVA_CLIENT_SECRET=%s\n", clientSecret)
	fmt.Printf("STRAVA_REFRESH_TOKEN=%s\n", tokenResponse["refresh_token"])

	// Offer to write to .env file
	fmt.Print("\nWould you like me to update the .env file automatically? (y/n): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response == "y" || response == "yes" {
		envContent := fmt.Sprintf(
			"STRAVA_CLIENT_ID=%s\nSTRAVA_CLIENT_SECRET=%s\nSTRAVA_REFRESH_TOKEN=%s\n",
			clientId,
			clientSecret,
			tokenResponse["refresh_token"],
		)

		err = os.WriteFile(".env", []byte(envContent), 0644)
		if err != nil {
			log.Fatal("Failed to write .env file: ", err)
		}

		fmt.Println("\n✓ .env file updated successfully!")
	}

	fmt.Println("\nYou can now run: make fetch_strava")
}
