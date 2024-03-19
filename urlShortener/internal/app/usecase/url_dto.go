package usecase


type CreateUrlReq struct {
	URL string `json:"url"`
}

type CreateUrlRes struct {
	ShortURL string `json:"short_url"`
}
type GetUrlReq struct {
	ShortURL string `json:"short_url"`
}

type GetUrlRes struct {
	URL string `json:"url"`
}