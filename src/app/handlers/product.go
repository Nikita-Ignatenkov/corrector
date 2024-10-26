package handlers

//
//import (
//	"corrector/repository"
//	"github.com/jackc/pgx/v5"
//	"github.com/labstack/echo/v4"
//	"github.com/pkg/errors"
//	"net/http"
//)
//
//func Product(db *pgx.Conn) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		id := c.Param("id")
//		product, err := repository.Product(db, id)
//		if err != nil {
//			return errors.Wrap(err, "can not get product")
//		}
//
//		data := product
//		return errors.Wrap(c.Render(http.StatusOK, "product.html", data), "can not render html")
//	}
//}
