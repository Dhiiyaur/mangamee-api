package bookmarkservice

import (
	"context"
	"mangamee-api/entity"
	bookmarkrepository "mangamee-api/repository/bookmark"
	"strings"

	"github.com/google/uuid"
)

type BookmarkServiceImpl struct {
	repository bookmarkrepository.BookmarkRepository
}

func NewBookmarkService(repository bookmarkrepository.BookmarkRepository) BookmarkService {
	return &BookmarkServiceImpl{
		repository: repository,
	}
}

func (s *BookmarkServiceImpl) GetBookmark(ctx context.Context, id string) (interface{}, error) {

	var bookmarkData entity.BookmarkRepository
	bookmarkData.Key = id
	result, err := s.repository.FindById(ctx, bookmarkData)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (s *BookmarkServiceImpl) InsertBookmark(ctx context.Context, mangaData interface{}) (string, error) {

	var bookmarkData entity.BookmarkRepository
	generateID := uuid.New()
	uuid := strings.Replace(generateID.String(), "-", "", -1)
	bookmarkData.Key = uuid
	bookmarkData.Data = mangaData

	err := s.repository.Insert(ctx, bookmarkData)
	if err != nil {
		return "", err
	}

	return bookmarkData.Key, nil
}
