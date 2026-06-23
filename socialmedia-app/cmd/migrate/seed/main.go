package main

import (
	"log"

	"github.com/olawuwo-abideen/go-project/socialmedia-app/internal/db"
	"github.com/olawuwo-abideen/go-project/socialmedia-app/internal/env"
	"github.com/olawuwo-abideen/go-project/socialmedia-app/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:password@localhost/socialmedia?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
