package usecase

import (
	"context"
	"testing"
	"urlShortener/urlShortener/internal/entity"
	mocks "urlShortener/urlShortener/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestURLUsecase_CreateShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockURLRepo(ctrl)
	usecase := New(mockRepo)

	ctx := context.Background()
	request := &CreateUrlReq{URL: "https://example.com"}
	expectedShortURL := "1"

	mockRepo.EXPECT().CreateShortURL(ctx, gomock.Any()).Return(&entity.URLShortener{ShortURL: expectedShortURL}, nil)

	response, err := usecase.CreateShortURL(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedShortURL, response.ShortURL)
}

func TestURLUsecase_GetShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
   
	mockRepo := mocks.NewMockURLRepo(ctrl)
	usecase := New(mockRepo)
   
	ctx := context.Background()
	request := &GetUrlReq{ShortURL: "1"}
	expectedURL := "https://example.com"
	expectedID := 1
   
	mockRepo.EXPECT().GetShortURL(ctx, expectedID).Return(&entity.URLShortener{URL: expectedURL}, nil)
   
	response, err := usecase.GetShortURL(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedURL, response.URL)
   }
