package adapters

import (
	"context"
	"urlShortener/urlShortener/internal/entity"
)
//go:generate mockgen -package internal -destination ../mocks/repository.go . URLRepo
type URLRepo interface {
	CreateShortURL(ctx context.Context, u *entity.URLShortener) (*entity.URLShortener, error)
	GetShortURL(ctx context.Context, id int) (*entity.URLShortener, error)
}
