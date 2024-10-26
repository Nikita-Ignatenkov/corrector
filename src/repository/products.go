package repository

import (
	"context"
	"corrector/models"

	"github.com/pkg/errors"
)

func ProductList(db PgxDB, query string) (models.ProductList, error) {
	var params []interface{}

	sql := "SELECT id, name, price, trademark_id FROM product"
	if query != "" {
		sql += " WHERE name ILIKE $1"

		params = append(params, "%"+query+"%")
	}

	rows, err := db.Query(context.Background(), sql, params...)

	if err != nil {
		return nil, errors.Wrap(err, "can not select from db")
	}

	res := make([]models.Product, 0)

	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.TrademarkID); err != nil {
			return nil, errors.Wrap(err, "can not scan product")
		}

		res = append(res, p)
	}

	return res, nil
}

func Product(db PgxDB, id string) (models.Product, error) {
	var p = models.Product{}
	err := db.QueryRow(context.Background(), "SELECT id, name, price, trademark_id FROM product WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price, &p.TrademarkID)
	if err != nil {
		return models.Product{}, errors.Wrap(err, "can not select from db")
	}
	return p, nil
}

const upsertProductQuery = `
	INSERT INTO product (id, name, price, trademark_id) 
    VALUES ($1, $2, $3, $4) ON CONFLICT (id) 
    DO UPDATE SET name=$2, price=$3, trademark_id=$4
`

func UpsertProduct(db PgxDB, product models.Product) error {
	if _, err := db.Exec(context.Background(), upsertProductQuery, product.ID, product.Name, product.Price,
		product.TrademarkID); err != nil {
		return errors.Wrapf(err, "can not upsert product %+v", product)
	}

	return nil
}

func UpsertProducts(db PgxDB, products []models.Product) error {
	for _, product := range products {
		if err := UpsertProduct(db, product); err != nil {
			return errors.Wrapf(err, "can not upsert products %+v", product)
		}
	}

	return nil
}
