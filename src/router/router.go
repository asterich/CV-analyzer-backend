package router

import (
	v1 "github.com/CV-analyzer-backend/src/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	// TODO: use a config file to set the mode
	gin.SetMode("debug")

	r := gin.Default()
	r.Static("/", "./static")

	api_v1 := r.Group("/api/v1")

	{
		api_v1.POST("/upload", v1.UploadFile)
		api_v1.POST("/upload/multi", v1.UploadMultiFile)
	}

	return r
}
