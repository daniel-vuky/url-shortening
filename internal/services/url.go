package services

import (
	"github.com/daniel-vuky/url-shortening/internal/config"
	"github.com/daniel-vuky/url-shortening/internal/models"
	"github.com/daniel-vuky/url-shortening/internal/storage/postgres"
	"github.com/daniel-vuky/url-shortening/internal/utils"
)

type IURL interface {
	CreateURL(requestParam *models.CreateURLRequest) (models.CreateURLResponse, error)
	GetURL(shortenCode string) (string, error)
	GetURLStats(shortenCode string) models.URLStats
	DeleteURL(userID string, shortenCode string) error
	GetListURLByUser(userID string, limit, offset int) ([]models.CreateURLResponse, error)
	IncreaseURLUsedCount(shortenCode string) error
}

type URLService struct {
	querier   postgres.Querier
	shortener utils.Shortener
	config    *config.Config
}

// NewService
// Init new service
func NewUrlService(querier postgres.Querier, shortener utils.Shortener, config *config.Config) *URLService {
	return &URLService{
		querier:   querier,
		shortener: shortener,
		config:    config,
	}
}

func (s *URLService) CreateURL(requestParam *models.CreateURLRequest) (models.CreateURLResponse, error) {
	return models.CreateURLResponse{}, nil
}

func (s *URLService) GetURL(shortenCode string) (string, error) {
	return "", nil
}

func (s *URLService) GetURLStats(shortenCode string) models.URLStats {
	return models.URLStats{}
}

func (s *URLService) DeleteURL(userID string, shortenCode string) error {
	return nil
}

func (s *URLService) GetListURLByUser(userID string, limit, offset int) ([]models.CreateURLResponse, error) {
	return []models.CreateURLResponse{}, nil
}

func (s *URLService) IncreaseURLUsedCount(shortenCode string) error {
	return nil
}
