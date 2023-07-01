package main

import (
	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/router"
)

func main() {
	model.InitDb()
	r := router.InitRouter()
	r.Run(":8080")
}
