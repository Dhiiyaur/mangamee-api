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
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DataChapters struct {
	ChapterName string  `json:"chapter_name"`
	Images      []Image `json:"images"`
}

type Image struct {
	Image string `json:"image"`
}

type MangaSource struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
