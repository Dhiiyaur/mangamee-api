package entity

type MangaData struct {
	Id             string       `json:"id,omitempty"`
	Cover          string       `json:"cover,omitempty"`
	Title          string       `json:"title,omitempty"`
	LastChapter    string       `json:"last_chapter,omitempty"`
	LastRead       string       `json:"last_read,omitempty"`
	Status         string       `json:"status,omitempty"`
	Summary        string       `json:"summary,omitempty"`
	Chapters       []Chapter    `json:"chapters,omitempty"`
	Images         DataChapters `json:"data_images,omitempty"`
	OriginalServer string       `json:"original_server,omitempty"`
}

type Chapter struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DataChapters struct {
	ChapterName string  `json:"chapter_name,omitempty"`
	Images      []Image `json:"images,omitempty"`
}

type Image struct {
	Image string `json:"image,omitempty"`
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

//

type MangaRepository struct {
	Key  string
	Data interface{}
}

type BookmarkRepository struct {
	Key  string
	Data interface{}
}

type ShortenerRepository struct {
	LongUrl string
	Key     string
}
