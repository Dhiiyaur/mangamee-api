package linkshortenerservice

import (
	"crypto/sha256"
	"fmt"
	"mangamee-api/src/entity"
	shortenerrepository "mangamee-api/src/repository/shortener"
	"math/big"

	"github.com/itchyny/base58-go"
	"golang.org/x/net/context"
)

type ShortenerServiceImpl struct {
	repository shortenerrepository.ShortenerRepository
}

func NewShortenerService(repository shortenerrepository.ShortenerRepository) ShortenerService {
	return &ShortenerServiceImpl{
		repository: repository,
	}
}

func (s *ShortenerServiceImpl) InsertLink(ctx context.Context, url string) (entity.ShortenerRepository, error) {

	var returnData entity.ShortenerRepository
	urlHashBytes := sha256Generator(url)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	if err != nil {
		return returnData, err
	}

	returnData.Key = finalString[:8]
	returnData.LongUrl = url

	err = s.repository.Insert(ctx, returnData)
	if err != nil {
		return returnData, err
	}
	return returnData, nil
}

func (s *ShortenerServiceImpl) GetLink(ctx context.Context, id string) (entity.ShortenerRepository, error) {

	var data entity.ShortenerRepository
	data.Key = id

	data, err := s.repository.FindById(ctx, data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func base58Encoded(bytes []byte) (string, error) {

	encoding := base58.BitcoinEncoding
	encode, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encode), nil
}

func sha256Generator(s string) []byte {

	alg := sha256.New()
	alg.Write([]byte(s))
	return alg.Sum(nil)
}
