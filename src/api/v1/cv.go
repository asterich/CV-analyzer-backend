package v1

import (
	"log"
	"os"
	"strconv"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 用于计算两个 CV 切片的交集
func intersectCVs(cv1, cv2 []model.CV) []model.CV {
	var intersectCV []model.CV
	cv2map := make(map[int]int, len(cv2))
	for _, v := range cv2 {
		cv2map[v.ID] = 1
	}
	for _, v := range cv1 {
		if _, ok := cv2map[v.ID]; ok {
			intersectCV = append(intersectCV, v)
		}
	}
	return intersectCV
}

func GetCV(c *gin.Context) {
	filename := c.Query("filename")
	id := c.Query("id")
	name := c.Query("name")
	degree := c.Query("degree")
	workingYears := c.Query("workingYears")
	age := c.Query("age")

	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	var returncv []model.CV

	if id != "" {
		GetCVById(c)
		return
	}

	if filename != "" {
		GetCVByFilename(c)
		return
	}

	if name != "" {
		date, err := model.GetCVsByName(name, pagesize, page)
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
		returncv = date
	}

	if degree != "" {
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
		if returncv != nil {
			returncv = intersectCVs(returncv, data)
		}
	}

	if workingYears != "" {
		workingYearsInt, _ := strconv.Atoi(workingYears)
		data, err := model.GetCVsGreaterThanWorkingYears(workingYearsInt, pagesize, page)
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
		if returncv != nil {
			returncv = intersectCVs(returncv, data)
			if len(returncv) == 0 {
				c.JSON(404, gin.H{
					"code": 404,
					"msg":  "CV not found",
				})
				return
			}
		} else {
			returncv = data
		}
	}

	if age != "" {
		ageInt, _ := strconv.Atoi(age)
		data, err := model.GetCVLesserThanAge(ageInt, pagesize, page)
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
		if returncv != nil {
			returncv = intersectCVs(returncv, data)
			if len(returncv) == 0 {
				c.JSON(404, gin.H{
					"code": 404,
					"msg":  "CV not found",
				})
				return
			}
		} else {
			returncv = data
		}
	}

	if name == "" && degree == "" && workingYears == "" && age == "" {
		GetAllCVs(c)
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": returncv,
	})
}

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

func GetCVsLesserThanAge(c *gin.Context) {
	ageStr := c.Query("age")
	age, _ := strconv.Atoi(ageStr)
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetCVLesserThanAge(age, pagesize, page)
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

func GetAllCVs(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, err := model.GetAllCVs(pagesize, page)
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
