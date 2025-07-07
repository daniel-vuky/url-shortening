package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/config"
	"github.com/daniel-vuky/url-shortening/internal/handlers"
	"github.com/daniel-vuky/url-shortening/internal/routes"
	"github.com/daniel-vuky/url-shortening/internal/services"
	"github.com/daniel-vuky/url-shortening/internal/storage/postgres"
	"github.com/daniel-vuky/url-shortening/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer(port int) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load the server config")
	}

	dbConnection := createDBConnection(config)
	defer dbConnection.Close()

	appHandlers := createHandler(dbConnection, config)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: routes.InitRoutes(appHandlers).GetMux(),
	}

	go func() {
		fmt.Println("Starting server on port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}

func createHandler(db *pgxpool.Pool, config *config.Config) *handlers.Handler {
	queries := postgres.New(db)
	shortener := utils.NewShortener()
	urlServices := services.NewUrlService(queries, *shortener, config)

	return handlers.NewHandler(urlServices)
}

func createDBConnection(config *config.Config) *pgxpool.Pool {
	dbUrl := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		config.Database.Type,
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.SSLMode,
	)
	ctx := context.Background()
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Can not parse db config")
	}
	dbConfig.MaxConns = config.Database.MaxConns
	dbConfig.MinConns = config.Database.MinConns
	dbConfig.HealthCheckPeriod = 30 * time.Second
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalf("Can not open connection to datbase")
	}

	return dbPool
}
