package v1

import (
	"log"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/gin-gonic/gin"
)

func GetAllDegree(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	log.Default().Println("pagesize: ", pagesize)
	log.Default().Println("page: ", page)
	degrees, err := model.GetAllDegree(pagesize, page)
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

func GetAllWorkingyears(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	log.Default().Println("pagesize: ", pagesize)
	log.Default().Println("page: ", page)
	workyears, err := model.GetAllWorkingyears(pagesize, page)
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
