package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naufalb95/trackr/internal/configs"
	"github.com/naufalb95/trackr/internal/routers"
)

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		fmt.Printf("Error when trying to load config: %v\n", err)
	}

	pool, err := pgxpool.New(context.Background(), config.GetDatabaseURL())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	} else {
		fmt.Printf("Connected to DB!")
	}

	defer pool.Close()

	r := routers.SetupRouter(pool)
	_ = r.Run(config.GetServicePort())
}
