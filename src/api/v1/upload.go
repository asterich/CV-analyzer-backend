package v1

import (
	"fmt"
	"net/http"
	"path/filepath"

	"log"

	"github.com/CV-analyzer-backend/src/converter"
	"github.com/CV-analyzer-backend/src/model"
	"github.com/CV-analyzer-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件上传失败",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, utils.UploadPath+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件保存失败",
		})
		return
	}

	// TODO: convert the file to CV and save it to the database
	cv, err := converter.ConvertDoc(utils.UploadPath + filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件转换失败",
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

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
		"url":  utils.UploadPath + filename,
	})
}

func UploadMultiFile(c *gin.Context) {
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
		if err := c.SaveUploadedFile(file, utils.UploadPath + filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  fmt.Sprintf("文件%s保存失败", filename),
			})
			continue
		}
		// TODO: convert the file to CV and save it to the database
		cv, err := converter.ConvertDoc(utils.UploadPath + filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  fmt.Sprintf("文件%s转换失败", filename),
			})
			continue
		}

		if err := model.SetCV(&cv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "文件保存失败",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
	})
}
