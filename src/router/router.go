package router

import (
	"github.com/asterich/CV-analyzer-backend/middleware"
	v1 "github.com/asterich/CV-analyzer-backend/src/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	// TODO: use a config file to set the mode
	gin.SetMode("debug")

	r := gin.Default()
	r.Static("/static", "./static")

	api_v1 := r.Group("/api/v1")

	{
		// upload file
		api_v1.POST("/cv", v1.UploadCV)
		api_v1.POST("/cv/multi", v1.UploadMultiCV)
		api_v1.POST("/positions", v1.UploadPosition)

		api_v1.GET("/cv", v1.GetCV)
		// cv
		api_v1.DELETE("/cv/filename/:filename", v1.DeleteCVByFilename)
		api_v1.DELETE("/cv/id/:id", v1.DeleteCVByID)

		// position
		api_v1.GET("/positions", v1.GetPositions)
		api_v1.DELETE("/positions/id/:id", v1.DeletePosition)

		//for visualize
		api_v1.GET("/count_degree", v1.GetCountDegree)
		api_v1.GET("/count_workingyears", v1.GetCountWorkingyears)
	}

	r.Use(middleware.CORS())

	return r
}
