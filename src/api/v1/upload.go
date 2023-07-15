package v1

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"log"

	"github.com/asterich/CV-analyzer-backend/src/converter"
	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func UploadCV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	log.Printf("filename: %s\n", utils.UploadPath+"cv/"+filename)
	if err := c.SaveUploadedFile(file, utils.UploadPath+"cv/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		return
	}

	absFilename, _ := filepath.Abs(utils.UploadPath + "cv/" + filename)
	cv, err := converter.ConvertDocToCV(absFilename)
	if err != nil {
		log.Printf("Error converting file %s, err = %s\n", utils.UploadPath+"cv/"+filename, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		err = os.Remove(utils.UploadPath + "cv/" + filename)
		if err != nil {
			log.Printf("Error removing file %s, err = %s\n", utils.UploadPath+"cv/"+filename, err.Error())
		}
		return
	}

	cv.Filename = utils.UploadPath + "cv/" + filename
	if err := model.SetCV(&cv); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Internal server error",
		})
		err = os.Remove(utils.UploadPath + "cv/" + filename)
		if err != nil {
			log.Printf("Error removing file %s, err = %s\n", utils.UploadPath+"cv/"+cv.Filename, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
		"url":  utils.UploadPath + filename,
	})
}

func UploadMultiCV(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
		return
	}

	files := form.File["files[]"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		log.Printf("filename: %s\n", filename)
		if err := c.SaveUploadedFile(file, utils.UploadPath+"cv/"+filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  fmt.Sprintf("文件%s保存失败", filename),
			})
			return
		}
		// TODO: convert the file to CV and save it to the database
		cv, err := converter.ConvertDocToCV(utils.UploadPath + "cv/" + filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  fmt.Sprintf("文件%s转换失败", filename),
			})
			return
		}

		if err := model.SetCV(&cv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "文件保存失败",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
	})
}

func UploadPosition(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, utils.UploadPath+"position/"+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件保存失败",
		})
		return
	}

	positions, err := converter.ConvertDocToPositions(utils.UploadPath + "position/" + filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件转换失败",
		})
		return
	}

	for _, position := range positions {
		if err := model.CreatePosition(&position); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "岗位信息创建失败",
			})
			return
		}
	}
}
