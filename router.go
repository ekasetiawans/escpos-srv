package main

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func init() {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/printers", func(ctx *gin.Context) {
		ctx.JSON(200, getPrinters())
	})

	router.POST("/print", func(ctx *gin.Context) {

		data := &struct {
			Printer string                `form:"printer" binding:"required"`
			File    *multipart.FileHeader `form:"file" binding:"required"`
		}{}

		err := ctx.Bind(data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		executable, _ := os.Executable()
		executableDirectory := filepath.Dir(executable)
		printJobDirectory := executableDirectory + "/jobs/" + data.Printer
		if _, err := os.Stat(printJobDirectory); os.IsNotExist(err) {
			os.Mkdir(printJobDirectory, 0777)
		}

		temporaryFile := printJobDirectory + "/" + data.File.Filename
		err = ctx.SaveUploadedFile(data.File, temporaryFile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		err = Print(temporaryFile, data.Printer)
		if err == nil {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":    true,
			"messages": err.Error(),
		})
	})
}
