package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	result := gin.Default()
	result.GET("/ping", func(content *gin.Context) {
		content.JSON(200, gin.H{
			"message": "haoxuan",
		})
	})
	err := result.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
