package gohttp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	"urlShortener/urlShortener/internal/app/usecase"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error during reading body: "+err.Error(), http.StatusBadRequest)
		return
	}
	url := string(b)
	//var req usecase.CreateUserReq
	req := usecase.CreateUrlReq{
		URL: url,
	}
	res, err := h.service.CreateShortURL(ctx, &req)
	if err != nil {
		http.Error(w, "can't create short url: "+err.Error(), http.StatusInternalServerError)
		return
	}
	shortenedURL := fmt.Sprintf("http://%s/%s", r.Host, res.ShortURL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortenedURL))

}

func (h *Handler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	shortURL := chi.URLParam(r, "shortURL")

	req := usecase.GetUrlReq{
		ShortURL: shortURL,
	}
	res, err := h.service.GetShortURL(ctx, &req)
	if err != nil {
		http.Error(w, "can't get original url: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.URL))

}
