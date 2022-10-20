package entity

type RequestParams struct {
	Source     string
	PageNumber string
	MangaId    string
	ChapterId  string
	Search     string
	ImageProxy string
	Path       string
}

type RequestLink struct {
	Url string
}
