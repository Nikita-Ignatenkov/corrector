package app

import (
	"corrector/app/parser"
	"corrector/check"
	"log"
)

func (s *Server) Import() {
	log.Println("Starting the import of all products...")

	products, err := parser.ParseAllProducts()
	if err != nil {
		log.Printf("Cannot read API response: %v\n", err)
		return
	}

	if err = check.UpsertProducts(s.DB, products); err != nil {
		log.Printf("Cannot insert products: %v\n", err)
		return
	}

	log.Printf("All products imported successfully.")
}
