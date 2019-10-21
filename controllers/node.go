package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianggushi/topstory/models"
)

// ListNodes .
func ListNodes(c *gin.Context) {
	nodes, err := models.GetNodes()
	if err != nil {
		c.JSON(404, formatJSON(nil, err))
		return
	}
	c.JSON(200, formatJSON(nodes, nil))
}
