package localrepo

import (
	"context"
	"testing"
	"urlShortener/urlShortener/internal/entity"

	"github.com/stretchr/testify/suite"
)

type hashTestSuite struct {
	suite.Suite
	urlRepo *repository
}

func TestHashRun(t *testing.T) {
	suite.Run(t, new(hashTestSuite))
}

func (d *hashTestSuite) SetupTest() {
	d.urlRepo = New()

}

func (d *hashTestSuite) TestCreateShortURL() {
	urlShortener := &entity.URLShortener{URL: "https://example.com"}
	createdURL, err := d.urlRepo.CreateShortURL(context.Background(), urlShortener)

	d.NoError(err)
	d.NotNil(createdURL)

	duplicateURL, err := d.urlRepo.CreateShortURL(context.Background(), urlShortener)
	d.Error(err)
	d.Nil(duplicateURL)
	d.Equal("url: "+urlShortener.URL+" - already exist", err.Error())
}

func (d *hashTestSuite) TestGet() {
	testURL := "https://example.com"
	d.urlRepo.CreateShortURL(context.Background(), &entity.URLShortener{URL: testURL})

	shortURL, err := d.urlRepo.GetShortURL(context.Background(), 1)
	d.NoError(err)
	d.NotNil(shortURL)
	d.Equal(testURL, shortURL.URL)

	nonExistentShortURL, err := d.urlRepo.GetShortURL(context.Background(), 999)
	d.Error(err)
	d.Nil(nonExistentShortURL)
	d.EqualError(err, "not found")
}
