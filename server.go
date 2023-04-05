package main

import (
	"log"
	"os"

	"github.com/PeemXD/expenses-gin/auth"
	"github.com/PeemXD/expenses-gin/database"
	"github.com/PeemXD/expenses-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	handler := handler.NewHandler(database.Db)
	r := gin.New()

	r.POST("/login", auth.LoginHandler)

	protected := r.Group("/", auth.Authorization)

	protected.POST("/expenses", handler.CreateExpensesHandler)

	protected.GET("/expenses/:id", handler.GetsExpensesHandler)

	protected.PUT("/expenses/:id", handler.EditExpensesHandler)

	protected.GET("/expenses", handler.GetExpensesHandler)

	log.Fatal(r.Run(os.Getenv("PORT")))

}
