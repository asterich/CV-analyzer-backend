package v1

import (
	"os"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCVById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, err := model.GetCVById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func GetCVByFilename(c *gin.Context) {
	path := c.Query("filename")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetCVByFilename(path, pagesize, page)
	if err != nil {
		// TODO: Avoid hard coding
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "CV not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Internal server error",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": []model.CV{data},
	})
}

func GetCVsByName(c *gin.Context) {
	name := c.Query("name")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetCVsByName(name, pagesize, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "CV not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Internal server error",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func GetCVsByDegree(c *gin.Context) {
	degree := c.Query("degree")
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetCVsByDegree(degree, pagesize, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "CV not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Internal server error",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func GetCVsGreaterThanWorkingYears(c *gin.Context) {
	workingYearsStr := c.Query("workingYears")
	workingYears, _ := strconv.Atoi(workingYearsStr)
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetCVsGreaterThanWorkingYears(workingYears, pagesize, page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "CV not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Internal server error",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func DeleteCVByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	cv, err := model.GetCVById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "CV not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Internal server error",
			})
			return
		}
	}
	err = model.DeleteCVByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	if err = os.Remove(utils.UploadPath + cv.Filename); err != nil {
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

func DeleteCVByFilename(c *gin.Context) {
	filename := c.Param("filename")
	err := model.DeleteCVByFilename(filename)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	if err = os.Remove(utils.UploadPath + filename); err != nil {
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

// TODO: Add more APIs
// There is no need to add APIs for creating
