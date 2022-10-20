package entity

type MangaRepository struct {
	Key  string
	Data interface{}
}

type BookmarkRepository struct {
	Key  string
	Data interface{}
}

type LinkRepository struct {
	LongUrl string
	Key     string
}
