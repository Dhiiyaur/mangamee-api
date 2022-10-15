package mangaservice

import (
	"context"
	"errors"
	"mangamee-api/src/config"
	"mangamee-api/src/entity"
	"mangamee-api/src/exception"
	"mangamee-api/src/helper"
	mangarepository "mangamee-api/src/repository/manga"
)

type MangaServiceImpl struct {
	repository mangarepository.MangaRepository
}

func NewMangaService(repository mangarepository.MangaRepository) MangaService {
	return &MangaServiceImpl{
		repository: repository,
	}
}

func (s *MangaServiceImpl) GetIndex(ctx context.Context, params entity.MangaParams) ([]entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToArrayResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadIndex(params)
		if err != nil {
			return MangaData, err
		}
		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := MangatownIndex(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":
		MangaData, err := MangabatIndex(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":
		MangaData, err := MaidmyIndex(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	return nil, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}

func (s *MangaServiceImpl) GetSearch(ctx context.Context, params entity.MangaParams) ([]entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToArrayResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := MangareadSearch(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "2":

		MangaData, err := MangatownSearch(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":

		MangaData, err := MangabatSearch(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err := MaidmySearch(params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	return nil, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}

func (s *MangaServiceImpl) GetDetail(ctx context.Context, params entity.MangaParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = MangareadDetail(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "2":

		MangaData, err = MangatownDetail(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":

		MangaData, err = MangabatDetail(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err = MaidmyDetail(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	return MangaData, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}

func (s *MangaServiceImpl) GetImage(ctx context.Context, params entity.MangaParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = MangareadImage(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "2":

		MangaData, err = MangatownImage(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = MangabatImage(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "4":

		MangaData, err = MaidmyImage(params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil
	}

	return MangaData, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}

func (s *MangaServiceImpl) GetChapter(ctx context.Context, params entity.MangaParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = MangareadChapter(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err = MangatownChapter(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = MangabatChapter(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "4":

		MangaData, err = MaidmyChapter(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil
	}

	return MangaData, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}

func (s *MangaServiceImpl) GetMeta(ctx context.Context, params entity.MangaParams) (entity.MangaData, error) {

	var MangaData entity.MangaData
	var err error

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv(config.Cfg) {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err = MangareadMetaTag(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err = MangatownMetaTag(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = MangabatMetaTag(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err = MaidmyMetaTag(params)
		if err != nil {
			return MangaData, err
		}

		if helper.IsProductionEnv(config.Cfg) {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil
	}

	return MangaData, exception.NewErrorMsg(
		exception.CodeBadRequest,
		errors.New("bad request"),
	)
}
