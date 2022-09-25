package entity

type MangaData struct {
	Id             string       `json:"id"`
	Cover          string       `json:"cover"`
	Title          string       `json:"title"`
	LastChapter    string       `json:"last_chapter"`
	LastRead       string       `json:"last_read"`
	Status         string       `json:"status"`
	Summary        string       `json:"summary"`
	Chapters       []Chapter    `json:"chapters"`
	Images         DataChapters `json:"data_images"`
	OriginalServer string       `json:"original_server"`
}

type Chapter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DataChapters struct {
	ChapterName string  `json:"chapter_name"`
	Images      []Image `json:"images"`
}

type Image struct {
	Image string `json:"image"`
}

type ReturnData struct {
	Data  MangaData
	Datas []MangaData
}

type MangaParams struct {
	Source     string
	PageNumber string
	MangaId    string
	ChapterId  string
	Search     string
	ImageProxy string
	Path       string
}

type MangaSource struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
