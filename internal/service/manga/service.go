package mangaservice

import (
	"errors"
	"fmt"
	"mangamee-api/internal/config"
	"mangamee-api/internal/entity"
	"mangamee-api/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: *repo,
	}
}

func (s *Service) GetIndex(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadIndex(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownIndex(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatIndex(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmyIndex(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}

	return nil, errors.New("bad request")
}

func (s *Service) GetSearch(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadSearch(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownSearch(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatSearch(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmySearch(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}
	return nil, errors.New("bad request")
}

func (s *Service) GetDetail(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadDetail(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownDetail(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatDetail(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmyDetail(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}
	return nil, errors.New("bad request")

}

func (s *Service) GetImage(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadImage(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownImage(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatImage(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmyImage(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}
	return nil, errors.New("bad request")
}

func (s *Service) GetChapter(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadChapter(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownChapter(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatChapter(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmyChapter(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}
	return nil, errors.New("bad request")
}

func (s *Service) GetMeta(params entity.MangaParams) (interface{}, error) {

	if config.Cfg.Server.Env == "PROD" {
		MangaData, err := s.repo.GetCache(params)
		if err == nil {
			return MangaData, nil
		}
	}

	fmt.Println("meta rag")

	switch params.Source {
	case "1":

		MangaData, err := MangareadMetaTag(params)

		fmt.Println(MangaData)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownMetaTag(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "3":

		MangaData, err := MangabatMetaTag(params)
		if err != nil {
			return nil, err
		}

		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil

	case "4":

		MangaData, err := MaidmyMetaTag(params)
		if err != nil {
			return nil, err
		}
		if config.Cfg.Server.Env == "PROD" {
			go s.repo.InsertCache(params, MangaData)
			go s.repo.InsertStatistic(params)
		}
		return MangaData, nil
	}

	return nil, errors.New("bad request")

}
