package main

import (
	"database/sql"

	"github.com/filipe/exagonal/adapters/db"
	"github.com/filipe/exagonal/application"
)

func main() {
	database, _ := sql.Open("sqlite3", "db.sqlite3")
	productDb := db.NewProductDb(database)
	productService := application.NewProductService(productDb)
	product, _ := productService.Create("Product 1", 10.0)
	productService.Enable(product)
}
