package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) CreateExpensesHandler(c *gin.Context) {
	var expense model.Expenses
	// var expenseForPg model.ExpensesForPg

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tagsArray := pq.Array(expense.Tags)

	if err := h.db.Raw(`
		INSERT INTO expenses (title, amount, note, tags)
		VALUES (?, ?, ?, ?)
		RETURNING id`,
		expense.Title, expense.Amount, expense.Note, tagsArray).Scan(&expense.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &expense)
}
