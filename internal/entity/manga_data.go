package entity

type MangaData struct {
	Id          string    `json:"Id"`
	Cover       string    `json:"Cover"`
	Title       string    `json:"Title"`
	LastChapter string    `json:"LastChapter"`
	LastRead    string    `json:"LastRead"`
	Status      string    `json:"Status"`
	Summary     string    `json:"Summary"`
	Chapters    []Chapter `json:"Chapters"`
	Images      []Image   `json:"Images"`
}

type Chapter struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type Image struct {
	Image string `json:"Image"`
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
