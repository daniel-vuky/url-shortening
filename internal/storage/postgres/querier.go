package postgres

import (
	"context"

	"github.com/daniel-vuky/url-shortening/internal/models"
)

type Querier interface {
	CreateURL(ctx context.Context, arg CreateURLParams) (models.URL, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (models.User, error)
	CreateVerifyEmailToken(ctx context.Context, arg CreateVerifyEmailTokenParams) (models.VerifyEmail, error)
	DeleteURL(ctx context.Context, id int32) (models.URL, error)
	GetListURL(ctx context.Context, arg GetListURLParams) ([]models.URL, error)
	GetURL(ctx context.Context, id int32) (models.URL, error)
	GetUser(ctx context.Context, email string) (models.User, error)
	UpdateURL(ctx context.Context, arg UpdateURLParams) (models.URL, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (models.User, error)
	UpdateVerifyEmailToken(ctx context.Context, arg UpdateVerifyEmailTokenParams) (models.VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
