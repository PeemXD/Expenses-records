package main

import (
	"github.com/PeemXD/expenses-gin/database"
	"github.com/PeemXD/expenses-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	handler := handler.NewHandler(database.Db)
	r := gin.New()

	r.POST("/login", loginHandler)

	protected := r.Group("/", authorization)

	protected.POST("/expenses", handler.CreateExpensesHandler)

	protected.GET("/expenses/:id", handler.GetsExpensesHandler)

	protected.PUT("/expemse/:id", handler.EditExpensesHandler)

	protected.GET("/expenses", handler.GetExpensesHandler)

	r.Run(":2565")
}
