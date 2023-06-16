package main

import (
	"encoding/base64"
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
			Printer string `json:"printer"`
			Data    string `json:"data"`
		}{}

		err := ctx.Bind(data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		if printer, ok := printers[data.Printer]; ok {
			//decode base64
			payload, err := base64.StdEncoding.DecodeString(data.Data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":    true,
					"messages": "Invalid data",
				})
				return
			}

			printer.Pool.AddJob(&PrintJob{
				Printer: data.Printer,
				Data:    payload,
			})

			ctx.JSON(200, gin.H{
				"message": "OK",
			})

			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":    true,
			"messages": "Printer not found",
		})
	})
}
