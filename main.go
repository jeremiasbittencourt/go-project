package main

import (
	"database/sql"
	"fmt"
	"github.com/jeremiasbittencourt/go-project/internal/infra/database"
	"github.com/jeremiasbittencourt/go-project/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID: "1234",
		Price: 10.0,
		Tax: 1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}