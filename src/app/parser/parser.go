package parser

import (
	"bytes"
	"corrector/apiToken"
	"corrector/models"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const itemsURL = "https://www.sima-land.ru/api/v5/item?p=100"

func getItemData(token string) ([]models.Product, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("accept", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		Get(itemsURL)

	if err != nil {
		return nil, fmt.Errorf("Error fetching item data: %w", err)
	}

	var items []models.Product
	err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&items)
	if err != nil {
		return nil, fmt.Errorf("Error decoding item data: %w", err)
	}

	return items, nil
}

func ParseAllProducts() (models.ProductList, error) {
	token, err := apiToken.GetJWTToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get JWT token")
	}

	p, err := getItemData(token)
	fmt.Println(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get item data")
	}

	return p, nil
}
