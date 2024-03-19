package pgrepo

import (
	"context"
	"fmt"
	"urlShortener/urlShortener/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepo interface {
	CreateShortURL(ctx context.Context, u *entity.URLShortener) (*entity.URLShortener, error)
	GetShortURL(ctx context.Context, id int) (*entity.URLShortener, error)
}

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Repository struct {
	db *pgxpool.Pool
}

func New(pg *pgxpool.Pool) *Repository {
	return &Repository{pg}
}

func (r *Repository) CreateShortURL(ctx context.Context, u *entity.URLShortener) (*entity.URLShortener, error) {
	var id int
	err := r.db.QueryRow(ctx, "INSERT INTO urls(url) VALUES($1) RETURNING id", u.URL).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't execute sql: %s", err.Error())
	}

	u.ShortURL = base62Encode(id)

	_, err = r.db.Exec(ctx, "UPDATE urls SET short_url = $1 WHERE id = $2", u.ShortURL, id)
	if err != nil {
		return nil, fmt.Errorf("can't execute sql: %s", err.Error())
	}
	return u, nil
}

func (r *Repository) GetShortURL(ctx context.Context, id int) (*entity.URLShortener, error) {
	row := r.db.QueryRow(ctx, "SELECT url FROM urls WHERE id=$1", id)
	url := &entity.URLShortener{}
	if err := row.Scan(&url.URL); err != nil {
		return nil, fmt.Errorf("can't scan url: %w", err)
	}

	return url, nil
}
func base62Encode(num int) string {
	encoded := ""
	for num > 0 {
		remainder := num % 62
		num /= 62
		encoded = string(base62Chars[remainder]) + encoded
	}
	return encoded
}
