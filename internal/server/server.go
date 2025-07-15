package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/daniel-vuky/url-shortening/internal/config"
	"github.com/daniel-vuky/url-shortening/internal/handlers"
	"github.com/daniel-vuky/url-shortening/internal/routes"
	"github.com/daniel-vuky/url-shortening/internal/services"
	"github.com/daniel-vuky/url-shortening/internal/storage/postgres"
	"github.com/daniel-vuky/url-shortening/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg    *config.Config
	server *http.Server
	db     *pgxpool.Pool
}

func NewServer(cfg *config.Config) *Server {
	dbConnection := createDBConnection(cfg)
	appHandlers := createHandler(dbConnection, cfg)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: routes.InitRoutes(appHandlers).GetMux(),
	}

	return &Server{
		cfg:    cfg,
		server: httpServer,
		db:     dbConnection,
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %d", s.cfg.Server.Port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.db.Close()
	return s.server.Shutdown(ctx)
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
	dbConfig.MaxConns = int32(config.Database.MaxConns)
	dbConfig.MinConns = int32(config.Database.MinConns)
	dbConfig.HealthCheckPeriod = 30 * time.Second
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil || dbPool.Ping(ctx) != nil {
		log.Fatalf("Can not open connection to datbase, %d", config.Database.MaxConns)
	}

	return dbPool
}
