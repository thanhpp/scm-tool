package main

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/thanhpp/scm/ent"
)

func main() {
	setupDB()
}

func setupDB() {
	client, err := ent.Open(
		"postgres",
		"host=localhost port=5432 user=user dbname=scm password=password sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}
