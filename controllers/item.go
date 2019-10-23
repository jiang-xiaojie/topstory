package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianggushi/topstory/models"
)

// GetItems get items
func GetItems(c *gin.Context) {
	var arg struct {
		ID int `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&arg)
	if err != nil {
		c.JSON(400, formatJSON(nil, err))
		return
	}
	items, err := models.GetItemsByNodeID(arg.ID)
	if err != nil {
		c.JSON(404, formatJSON(nil, err))
		return
	}
	c.JSON(200, formatJSON(items, nil))
}
