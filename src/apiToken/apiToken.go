package apiToken

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"os"
)

type AuthResponse struct {
	Token string `json:"token"`
}

func GetJWTToken() (string, error) {
	err := godotenv.Load("/home/nikita/Projects/gitlab.sima-land.ru/corrector/src/dev.env")
	if err != nil {
		return "", fmt.Errorf("Error loading .env file: %w", err)
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	phone := os.Getenv("PHONE")

	client := resty.New()
	resp, err := client.R().
		SetHeader("accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"email":      email,
			"password":   password,
			"phone":      phone,
			"regulation": true,
		}).
		Post("https://www.sima-land.ru/api/v5/signin")

	if err != nil {
		return "", fmt.Errorf("Error getting JWT token: %w", err)
	}

	var authResp AuthResponse
	err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&authResp)
	if err != nil {
		return "", fmt.Errorf("Error decoding JWT token: %w", err)
	}
	return authResp.Token[7:], nil
}
