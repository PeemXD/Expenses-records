package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) EditExpensesHandler(c *gin.Context) {

	var expense model.Expenses
	id := c.Param("id")

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// check id existed
	if result := h.db.Raw("SELECT id FROM expenses WHERE id = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
	}

	// update base on primarykey
	if result := h.db.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, &expense)
}
