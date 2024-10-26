// parser.go

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

//func decodeCategoryResponse(body io.Reader) (models.ProductList, error) {
//	var items struct {
//		Items []struct {
//			models.Product
//		}
//	}
//
//	err := json.NewDecoder(body).Decode(&items)
//	if err != nil {
//		return nil, errors.Wrap(err, "cannot decode response")
//	}
//
//	p := models.ProductList{}
//
//	return p, nil
//}

const itemsURL = "https://www.sima-land.ru/api/v5/item?p=50"

// Функция для получения товаров
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

func ParseCategory(categoryID int) (models.ProductList, error) {
	token, err := apiToken.GetJWTToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get JWT token")
	}

	// Получаем данные о товарах
	p, err := getItemData(token)
	fmt.Println(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get item data")
	}

	// способ извлечь информацию из items
	// Например, если items - это []models.Product:
	//body := bytes.NewReader([]byte{}) // Это временное значение
	//
	//// используем items для обработки
	//p, t, err := decodeCategoryResponse(body)
	//if err != nil {
	//	return nil, nil, err
	//}

	return p, nil
}
