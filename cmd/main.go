package main

import (
	"awesomeProject3/internal/biz"
	"awesomeProject3/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, 1)
	})
	lh := handler.LogHandler{biz.LS}
	r.POST("/log", lh.CreateLog)
	r.DELETE("/log/:id", lh.DeleteLog)
	r.PUT("/log", lh.UpdateLog)
	r.GET("/log", lh.QueryLog)
	r.GET("/log/:id", lh.ViewLog)
	r.Run()
}
