package v1

import (
	"os"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/utils"
	"github.com/gin-gonic/gin"
)

func GetAllPositions(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetAllPositions(pagesize, page)
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

func GetPositionsByName(c *gin.Context) {
	name := c.Query("name")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
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

func GetPositionsByCompany(c *gin.Context) {
	company := c.Query("company")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetPositionsByCompany(company, pagesize, page)
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

func GetPositionsByDepartment(c *gin.Context) {
	department := c.Query("department")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetPositionsByDepartment(department, pagesize, page)
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

func GetPositionsByMajor(c *gin.Context) {
	major := c.Query("major")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetPositionsByMajor(major, pagesize, page)
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

func GetPositionsByDegree(c *gin.Context) {
	degree := c.Query("degree")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetPositionsByDegree(degree, pagesize, page)
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

func GetPositionsByWorkingYears(c *gin.Context) {
	workingYears, _ := strconv.Atoi(c.Query("working_years"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	positions, err := model.GetPositionsGreaterThanWorkingYears(workingYears, pagesize, page)
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

func DeletePosition(c *gin.Context) {
	filename := c.Param("filename")
	err := model.DeletePositionByFilename(filename)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	file, err := os.Open(utils.UploadPath + filename)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}
	defer file.Close()

	if err = os.Remove(filename); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
