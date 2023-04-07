package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) EditExpensesHandler(c *gin.Context) {

	id := c.Param("id")

	// check id existed
	var existedExpense model.Expenses
	if result := h.db.Where("id = ?", id).First(&existedExpense); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var expense model.Expenses
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update base on primarykey
	expense.ID = existedExpense.ID
	if result := h.db.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &expense)
}
