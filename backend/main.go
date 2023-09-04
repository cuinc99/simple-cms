package main

import (
	"github.com/cuinc99/simple-cms/controllers/postcontroller"
	"github.com/cuinc99/simple-cms/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/api/posts", postcontroller.Index)
	r.GET("/api/post/:id", postcontroller.Show)
	r.POST("/api/post", postcontroller.Create)
	r.PUT("/api/post/:id", postcontroller.Update)
	r.DELETE("/api/post", postcontroller.Delete)

	r.Run("127.0.0.1:4444")
}
