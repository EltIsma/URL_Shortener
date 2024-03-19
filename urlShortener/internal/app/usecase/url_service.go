package usecase

import (
	"context"
	"math"
	"strings"
	"urlShortener/urlShortener/internal/adapters/pgrepo"

	"urlShortener/urlShortener/internal/entity"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type URLUsecase struct {
	repo pgrepo.URLRepo
}

func New(ur pgrepo.URLRepo) *URLUsecase {
	return &URLUsecase{repo: ur}
}

func (uc *URLUsecase) CreateShortURL(ctx context.Context, m *CreateUrlReq) (*CreateUrlRes, error) {
	u := &entity.URLShortener{
		URL: m.URL,
	}

	r, err := uc.repo.CreateShortURL(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUrlRes{
		ShortURL: r.ShortURL,
	}
	return res, nil
}

func (uc *URLUsecase) GetShortURL(ctx context.Context, m *GetUrlReq) (*GetUrlRes, error) {
	u := &entity.URLShortener{
		ShortURL: m.ShortURL,
	}
	id := base62Decode(u.ShortURL)
	//	fmt.Println(id)

	r, err := uc.repo.GetShortURL(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &GetUrlRes{
		URL: r.URL,
	}
	return res, nil

}

func base62Decode(str string) int {
	decoded := 0
	for i := 0; i < len(str); i++ {
		pos := strings.Index(base62Chars, string(str[i]))
		decoded += pos * int(math.Pow(62, float64(i)))
	}
	return decoded
}
