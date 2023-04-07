package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetExpensesHandler(c *gin.Context) {

	id := c.Param("id")
	var expense model.Expenses

	if result := h.db.Find(&expense, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, &expense)
}

func (h *Handler) GetsExpensesHandler(c *gin.Context) {

	var expenses []model.Expenses

	if result := h.db.Raw(`SELECT * FROM expenses`).Scan(&expenses); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
