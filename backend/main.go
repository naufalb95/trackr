package main

import (
	"github.com/naufalb95/trackr/internal/routers"
)

func main() {
	r := routers.SetupRouter()
	_ = r.Run(":8080")
}
