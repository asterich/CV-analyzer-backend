package main

import (
	"github.com/CV-analyzer-backend/src/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
