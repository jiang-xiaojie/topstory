package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianggushi/topstory/models"
)

// GetLastItem get lastitems
func GetLastItem(c *gin.Context) {
	var arg struct {
		ID int `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&arg)
	if err != nil {
		c.JSON(400, formatJSON(nil, err))
		return
	}
	lastItem, err := models.GetLastItemByNodeID(arg.ID)
	if err != nil {
		c.JSON(404, formatJSON(nil, err))
		return
	}
	c.JSON(200, formatJSON(lastItem, nil))
}
