package gohttp

import (
	"urlShortener/urlShortener/internal/app/usecase"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *usecase.URLUsecase
}
func New(s *usecase.URLUsecase) *Handler{
	return &Handler{
		service: s,
	}
}

func NewRouter(h *Handler) *chi.Mux {
	mux := chi.NewRouter()

	mux.Route("/", func(r chi.Router) {
		r.Post("/", h.CreateShortURL)
		r.Get("/{shortURL}", h.GetOriginalURL)
	})

	return mux
}
