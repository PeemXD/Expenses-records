package handler

import (
	"net/http"
	"strings"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) GetExpensesHandler(c *gin.Context) {
	id := c.Param("id")
	var expense model.Expenses
	var expenseForPg model.ExpensesForPg

	if result := h.db.Raw("SELECT * FROM expenses WHERE id = ?", id).Scan(&expenseForPg); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	expense = model.Expenses{ID: expenseForPg.ID, Title: expenseForPg.Title, Amount: expenseForPg.Amount, Note: expenseForPg.Note}

	if result := h.db.Raw("SELECT tags FROM expenses WHERE id = ?", id).Scan(pq.Array(&expense.Tags)); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}
	// Convert the Postgres array of tags to a slice of strings
	//? ["{food,beverage}"]
	//? pick expense.Tags[0] --> "{food,beverage}" to cut {} to --> "food,beverage"
	//? split string with "," --> "food,beverage" to ["food", "beverage"] ([]string{"food", "beverage"})
	tagsString := strings.Trim(expense.Tags[0], "{}")
	tags := strings.Split(tagsString, ",")
	expense.Tags = tags

	c.JSON(http.StatusOK, &expense)
}

func (h *Handler) GetsExpensesHandler(c *gin.Context) {

}
