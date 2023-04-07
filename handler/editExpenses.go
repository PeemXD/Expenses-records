package handler

import (
	"net/http"
	"strconv"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) EditExpensesHandler(c *gin.Context) {

	var expense model.Expenses
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	// idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

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
	expense.ID = uint(idInt)
	if result := h.db.Raw(`UPDATE expenses 
						   SET title = ?, amount = ?, note = ?, tags = ? 
						   WHERE id = ?`,
		expense.Title, expense.Amount, expense.Note, pq.Array(&expense.Tags), expense.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, &expense)
}
