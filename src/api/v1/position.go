package v1

import (
	"log"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
	name := c.Query("name")
	major := c.Query("major")
	degree := c.Query("degree")
	workingYearsBegin := c.Query("working_years_begin")
	workingYearsEnd := c.Query("working_years_end")

	if name != "" {
		GetPositionsByName(c)
	} else if major != "" {
		GetPositionsByMajor(c)
	} else if degree != "" {
		GetPositionsByDegree(c)
	} else if workingYearsBegin != "" || workingYearsEnd != "" {
		GetPositionsByWorkingYears(c)
	} else {
		GetAllPositions(c)
	}
}

func GetAllPositions(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	log.Default().Println("pagesize: ", pagesize)
	log.Default().Println("page: ", page)
	positions, err := model.GetAllPositions(pagesize, page)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	log.Default().Println("positions: ", positions)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": positions,
	})
}

func GetPositionsByName(c *gin.Context) {
	name := c.Query("name")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	positions, err := model.GetPositionsByName(name, pagesize, page)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": positions,
	})
}

// func GetPositionsByCompany(c *gin.Context) {
// 	company := c.Query("company")
// 	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
// 	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
// 	positions, err := model.GetPositionsByCompany(company, pagesize, page)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"code": 500,
// 			"msg":  "Internal server error",
// 		})
// 	}

// 	c.JSON(200, gin.H{
// 		"code": 200,
// 		"msg":  "success",
// 		"data": positions,
// 	})
// }

// func GetPositionsByDepartment(c *gin.Context) {
// 	department := c.Query("department")
// 	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
// 	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
// 	positions, err := model.GetPositionsByDepartment(department, pagesize, page)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"code": 500,
// 			"msg":  "Internal server error",
// 		})
// 	}

// 	c.JSON(200, gin.H{
// 		"code": 200,
// 		"msg":  "success",
// 		"data": positions,
// 	})
// }

func GetPositionsByMajor(c *gin.Context) {
	major := c.Query("major")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	positions, err := model.GetPositionsByMajor(major, pagesize, page)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": positions,
	})
}

func GetPositionsByDegree(c *gin.Context) {
	degree := c.Query("degree")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	positions, err := model.GetPositionsByDegree(degree, pagesize, page)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": positions,
	})
}

func GetPositionsByWorkingYears(c *gin.Context) {
	var workingYears model.IntDuration
	workingYears.Begin, _ = strconv.Atoi(c.DefaultQuery("working_years_begin", "0"))
	workingYears.End, _ = strconv.Atoi(c.DefaultQuery("working_years_end", "999"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	positions, err := model.GetPositionsInRangeOfWorkingYears(workingYears, pagesize, page)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": positions,
	})
}

func DeletePosition(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.DeletePositionByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
