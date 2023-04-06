package handler

import (
	"net/http"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetExpensesHandler(c *gin.Context) {

}

func (h *Handler) GetsExpensesHandler(c *gin.Context) {
	var expensesForPg []model.ExpensesForPg

	if result := h.db.Raw(`SELECT * FROM expenses`).Scan(&expensesForPg); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, expensesForPg)
}
