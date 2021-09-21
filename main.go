package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

type Product struct {
	id int
	owner Owner
	name string
	tier string
}

type Owner struct {
	id int
	name string
	gender string
}

type Filter struct {
	Key string
	Value string
	Wildcard bool
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:dev@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//fmt.Println(filterProducts(conn))
	fmt.Println(filterProducts(conn, Filter{
		Key:      "o.name",
		Value:    "Jeremy",
		Wildcard: false,
	}))
}

func filterProducts(conn *pgx.Conn, filters ...Filter) ([]Product, error) {
	var products []Product
	var rows pgx.Rows
	query := "SELECT p.id, p.name, p.tier, o.name, o.id, o.gender FROM products p JOIN owners o ON p.owner_id = o.id"
	if len(filters) > 0 {
		query += " WHERE "
		for _, f := range filters {
			query += fmt.Sprintf("%s LIKE '%s%%'", f.Key, f.Value)
		}
		fmt.Println(query)
		rows, _ = conn.Query(context.Background(), query)
	} else {
		rows, _ = conn.Query(context.Background(), query)
	}

	for rows.Next() {
		var productID int
		var productName string
		var productTier string
		var ownerName string
		var ownerID int
		var ownerGender string

		err := rows.Scan(&productID, &productName, &productTier, &ownerName, &ownerID, &ownerGender)
		if err != nil {
			log.Fatalln(err)
		}
		products = append(products, Product{
			id:    productID,
			owner: Owner{
				id:     ownerID,
				name:   ownerName,
				gender: ownerGender,
			},
			name:  productName,
			tier:  productTier,
		})
	}

	return products, rows.Err()
}