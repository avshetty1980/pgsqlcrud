package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/avshetty1980/pgsqlcrud/app_config"
	"github.com/avshetty1980/pgsqlcrud/products"
	_ "github.com/lib/pq"
)

func main() {

	app_config.ReadConfig()

	db, err := sql.Open("postgres", app_config.Cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	queries := products.New(db)

	//List products

	// productList, err := queries.ListProducts(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, product := range productList {
	// 	fmt.Println(product.Name, product.Price)
	// }

	//Get product

	// product, err := queries.GetProduct(ctx, 111)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		fmt.Println("no product with ID", 111)
	// 		return
	// 	}
	// 	log.Fatal(err)
	// }

	//Create Product

	// product, err := queries.CreateProduct(ctx, products.CreateProductParams{
	// 	Name:      "Coffee Machine",
	// 	Price:     "122.33",
	// 	Available: sql.NullBool{Bool: true, Valid: true},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Delete product

	// err = queries.DeleteProduct(ctx, 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//aggregate total
	total, err := queries.TotalPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)

}
