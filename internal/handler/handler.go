package handler

import (
	"awesomeProject3/internal/biz"
	"awesomeProject3/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LogHandler struct {
	LS biz.LogService
}

func (lh *LogHandler) CreateLog(c *gin.Context) {
	var logDto dto.LogDto
	if err := c.ShouldBindJSON(&logDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lh.LS.CreateLog(logDto)
}

func (lh *LogHandler) DeleteLog(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	lh.LS.DeleteLog(i)
	c.String(http.StatusOK, "deleted %s", id)
}

func (lh *LogHandler) UpdateLog(c *gin.Context) {
	var logDto dto.LogDto
	if err := c.ShouldBindJSON(&logDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lh.LS.UpdateLog(logDto)
}

func (lh *LogHandler) QueryLog(c *gin.Context) {
	var p dto.LogQueryParam
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.AsciiJSON(200, lh.LS.QueryLog(p))
}

func (lh *LogHandler) ViewLog(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	c.AsciiJSON(200, lh.LS.ViewLog(i))
}
