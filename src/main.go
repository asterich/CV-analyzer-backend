package main

import (
	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/router"
	"github.com/asterich/CV-analyzer-backend/src/utils"
)

func main() {
	model.InitDb()
	r := router.InitRouter()
	r.Run(utils.HttpPort)
}
