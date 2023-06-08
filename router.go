package main

import (
	"net/http"

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
		ctx.JSON(200, printers)
	})

	router.POST("/print", func(ctx *gin.Context) {

		data := &struct {
			Data string `json:"data"`
		}{}

		err := ctx.Bind(data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		pool.AddJob(&PrintJob{
			Data: data.Data,
		})

		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
