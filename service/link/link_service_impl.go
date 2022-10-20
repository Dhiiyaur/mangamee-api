package linkservice

import (
	"crypto/sha256"
	"fmt"
	"mangamee-api/entity"
	linkrepository "mangamee-api/repository/link"
	"math/big"

	"github.com/itchyny/base58-go"
	"golang.org/x/net/context"
)

type LinkServiceImpl struct {
	repository linkrepository.LinkRepository
}

func NewLinkService(repository linkrepository.LinkRepository) LinkService {
	return &LinkServiceImpl{
		repository: repository,
	}
}

func (s *LinkServiceImpl) InsertLink(ctx context.Context, url string) (entity.LinkRepository, error) {

	var returnData entity.LinkRepository
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

func (s *LinkServiceImpl) GetLink(ctx context.Context, id string) (entity.LinkRepository, error) {

	var data entity.LinkRepository
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
