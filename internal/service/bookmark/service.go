package bookmarkservice

import (
	"mangamee-api/internal/repository"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repo repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: *repo,
	}
}

func (s *Service) GenerateCode(bookmark interface{}) (interface{}, error) {

	generateID := uuid.New()
	uuid := strings.Replace(generateID.String(), "-", "", -1)

	err := s.repo.InsertBookmark(uuid, bookmark)
	if err != nil {
		return nil, err
	}
	return uuid, nil
}

func (s *Service) GetBookmark(code string) (interface{}, error) {

	r, err := s.repo.GetBookmark(code)
	if err != nil {
		return nil, err
	}
	return r, nil
}
