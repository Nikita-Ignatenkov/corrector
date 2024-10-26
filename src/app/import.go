package app

import (
	"corrector/app/parser"
	"corrector/repository"
	"log"
)

const (
	toysCategoryID    = 687
	apparelCategoryID = 3515
	sportCategoryID   = 16
)

func (s *Server) Import() {
	categoryNames := map[int]string{
		toysCategoryID:    "Игрушки",
		apparelCategoryID: "Одежда",
		sportCategoryID:   "Спорт",
	}

	for id, name := range categoryNames {

		products, err := parser.ParseCategory(id)
		if err != nil {
			log.Printf("can not read api response: %v\n", err)

			continue
		}

		if err = repository.UpsertProducts(s.DB, products); err != nil {
			log.Printf("can not insert products: %v\n", err)

			continue
		}

		log.Printf("Category \"%v\" saved", name)
	}
}
