package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PeemXD/expenses-gin/auth"
	"github.com/PeemXD/expenses-gin/database"
	"github.com/PeemXD/expenses-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	handler := handler.NewHandler(database.Db)
	r := gin.New()

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}

	r.POST("/login", auth.LoginHandler)

	protected := r.Group("/", auth.Authorization)

	protected.POST("/expenses", handler.CreateExpensesHandler)

	protected.GET("/expenses/:id", handler.GetsExpensesHandler)

	protected.PUT("/expenses/:id", handler.EditExpensesHandler)

	protected.GET("/expenses", handler.GetExpensesHandler)

	// log.Fatal(r.Run(os.Getenv("PORT"))) //if want to use graceful shutdown. it must use goroutine to run server
	go func() {
		if err := r.Run(os.Getenv("PORT")); err != nil {
			log.Fatalf("Could not gracefully shutdown server: %v\n", err)
		}
	}()

	//TODO graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	<-shutdown
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown server: %v\n", err)
	}

	log.Println("Server shutdown completed.")
}
