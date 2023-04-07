package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateExpensesHandler(c *gin.Context) {
	var expense model.Expenses

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &expense)
}
