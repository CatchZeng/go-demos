package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 设置上传限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传失败: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传错误: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("文件 %s 上传成功。参数 name=%s，email=%s.", file.Filename, name, email))
	})

	router.POST("/multipleUpload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传失败: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("上传错误: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("上传成功 %d 文件。 参数 name=%s，email=%s.", len(files), name, email))
	})

	router.Run(":8080")
}
