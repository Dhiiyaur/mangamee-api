package mangaservice

import (
	"context"

	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"
	mangarepository "mangamee-api/repository/manga"
	scrapper "mangamee-api/service/manga/scrapper"

	"go.uber.org/zap"
)

type MangaServiceImpl struct {
	repository mangarepository.MangaRepository
}

func NewMangaService(repository mangarepository.MangaRepository) MangaService {
	return &MangaServiceImpl{
		repository: repository,
	}
}

func (s *MangaServiceImpl) GetIndex(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToArrayResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := scrapper.MangareadIndex(ctx, params)
		if err != nil {
			return MangaData, err
		}
		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err := scrapper.MangatownIndex(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":
		MangaData, err := scrapper.MangabatIndex(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":
		MangaData, err := scrapper.MaidmyIndex(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	logger.Info("GetIndex source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return nil, exception.NewBadRequest("source not found")
}

func (s *MangaServiceImpl) GetSearch(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToArrayResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err := scrapper.MangareadSearch(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "2":

		MangaData, err := scrapper.MangatownSearch(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":

		MangaData, err := scrapper.MangabatSearch(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err := scrapper.MaidmySearch(ctx, params)
		if err != nil {
			return nil, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	logger.Info("GetSearch source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return nil, exception.NewBadRequest("source not found")
}

func (s *MangaServiceImpl) GetDetail(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = scrapper.MangareadDetail(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err = scrapper.MangatownDetail(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "3":

		MangaData, err = scrapper.MangabatDetail(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err = scrapper.MaidmyDetail(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	logger.Info("GetDetail source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return MangaData, exception.NewBadRequest("source not found")
}

func (s *MangaServiceImpl) GetImage(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = scrapper.MangareadImage(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "2":

		MangaData, err = scrapper.MangatownImage(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = scrapper.MangabatImage(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "4":

		MangaData, err = scrapper.MaidmyImage(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil
	}

	logger.Info("GetImage source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return MangaData, exception.NewBadRequest("source not found")
}

func (s *MangaServiceImpl) GetChapter(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	var MangaData entity.MangaData
	var err error

	switch params.Source {
	case "1":

		MangaData, err = scrapper.MangareadChapter(ctx, params)
		if err != nil {
			return MangaData, err
		}
		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err = scrapper.MangatownChapter(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = scrapper.MangabatChapter(ctx, params)
		if err != nil {
			return MangaData, err
		}
		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "4":

		MangaData, err = scrapper.MaidmyChapter(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil
	}

	logger.Info("GetChapter source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return MangaData, exception.NewBadRequest("source not found")
}

func (s *MangaServiceImpl) GetMeta(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var MangaData entity.MangaData
	var err error

	var mangaRepository entity.MangaRepository
	mangaRepository.Key = helper.GenerateKey(params)

	if helper.IsProductionEnv() {
		mangaData, err := s.repository.FindByIdCache(ctx, mangaRepository, params.Path)
		if err == nil {
			return helper.RepositoryToResponseData(mangaData), nil
		}
	}

	switch params.Source {
	case "1":

		MangaData, err = scrapper.MangareadMetaTag(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "2":

		MangaData, err = scrapper.MangatownMetaTag(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}
		return MangaData, nil

	case "3":

		MangaData, err = scrapper.MangabatMetaTag(ctx, params)
		if err != nil {
			return MangaData, err
		}
		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil

	case "4":

		MangaData, err = scrapper.MaidmyMetaTag(ctx, params)
		if err != nil {
			return MangaData, err
		}

		mangaRepository.Data = MangaData
		if helper.IsProductionEnv() {
			go s.repository.InsertCache(ctx, mangaRepository)
			go s.repository.InsertStatistic(ctx, params)
		}

		return MangaData, nil
	}

	logger.Info("GetMeta source not found", zap.Any("request_id", helper.GetRequestId(ctx)))
	return MangaData, exception.NewBadRequest("source not found")
}
