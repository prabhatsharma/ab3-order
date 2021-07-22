package main

import (
	"github.com/gin-gonic/gin"
	orderapi "github.com/prabhatsharma/ab3/api"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/order", orderapi.Get)
	r.POST("/order", orderapi.Post)

	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
