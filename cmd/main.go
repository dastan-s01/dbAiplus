package main

import (
	"context"
	"dbAiplus/db"
	"dbAiplus/internal/app/di"
	"dbAiplus/internal/app/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbName, sslMode)
	conn, err := db.ConnectionDB(dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer conn.Close()

	dependencies := di.NewDI(conn)

	router := gin.Default()
	handler := handlers.NewHandler(dependencies)
	handler.InitRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Сервер запущен на порту 8080")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Выключение сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка остановки сервера: %v", err)
	}

}
