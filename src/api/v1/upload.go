package v1

import (
	"net/http"
	"path/filepath"

	"log"

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
	if err := c.SaveUploadedFile(file, "upload/"+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件保存失败",
		})
		return
	}

	// TODO: convert the file to CV and save it to the database

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
		"url":  "upload/" + filename,
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
		if err := c.SaveUploadedFile(file, "upload/"+filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "文件保存失败",
			})
			return
		}
	}

	// TODO: convert the file to CV and save it to the database

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功",
	})
}
