package v1

import (
	"log"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/gin-gonic/gin"
)

func GetCountDegree(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))
	log.Default().Println("pagesize: ", pagesize)
	degrees, err := model.GetCountDegree(pagesize)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	log.Default().Println("degrees: ", degrees)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": degrees,
	})
}

func GetCountWorkingyears(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))
	log.Default().Println("pagesize: ", pagesize)
	workyears, err := model.GetCountWorkingyears(pagesize)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	log.Default().Println("workingyears: ", workyears)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": workyears,
	})
}
