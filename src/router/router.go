package router

import (
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
		api_v1.POST("/cv/upload", v1.UploadCV)
		api_v1.POST("/cv/upload/multi", v1.UploadMultiCV)
		api_v1.POST("/positions/upload", v1.UploadPosition)

		// cv
		api_v1.GET("/cv/filename", v1.GetCVByFilename)
		api_v1.GET("/cv/name", v1.GetCVsByName)
		api_v1.GET("/cv/degree", v1.GetCVsByDegree)
		api_v1.GET("/cv/working_years", v1.GetCVsGreaterThanWorkingYears)
		api_v1.DELETE("/cv/filename/:filename", v1.DeleteCVByFilename)

		// position
		api_v1.GET("/all_positions", v1.GetAllPositions)
		api_v1.GET("/positions/name", v1.GetPositionsByName)
		// api_v1.GET("/positions/company", v1.GetPositionsByCompany)
		// api_v1.GET("/positions/department", v1.GetPositionsByDepartment)
		api_v1.GET("/positions/degree", v1.GetPositionsByDegree)
		api_v1.GET("/positions/major", v1.GetPositionsByMajor)
		api_v1.GET("/positions/working_years", v1.GetPositionsByWorkingYears)
		api_v1.DELETE("/positions/filename/:filename", v1.DeletePosition)
	}

	return r
}
