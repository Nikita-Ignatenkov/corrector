package check

import (
	"context"
	"corrector/models"

	"github.com/pkg/errors"
)

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
