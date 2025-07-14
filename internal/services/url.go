package services

import (
	"context"
	"time"

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

// CreateURL
// Receives url param and start to create the shortening
func (s *URLService) CreateURL(requestParam *models.CreateURLRequest) (models.CreateURLResponse, error) {
	var expiresAt time.Time
	if requestParam.ExpiresAt != nil {
		expiresAt = *requestParam.ExpiresAt
	} else {
		expiresAt = time.Now().Add(365 * 24 * time.Hour) // hoặc giá trị mặc định bạn muốn
	}
	createUrlParams := postgres.CreateURLParams{
		OriginalUrl: requestParam.OriginalURL,
		ShortCode:   requestParam.ShortCode,
		ExpiresAt:   expiresAt,
	}
	urlCreated, err := s.querier.CreateURL(context.Background(), createUrlParams)
	if err != nil {
		return models.CreateURLResponse{}, err
	}
	return models.CreateURLResponse{
		ShortURL:    urlCreated.ShortCode,
		ShortCode:   urlCreated.ShortCode,
		OriginalURL: urlCreated.OriginalURL,
		ExpiresAt:   &urlCreated.ExpiresAt.Time,
	}, nil
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
