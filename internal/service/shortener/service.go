package linkshortenerservice

import (
	"crypto/sha256"
	"fmt"
	"mangamee-api/internal/repository"
	"math/big"

	"github.com/itchyny/base58-go"
)

type Service struct {
	repo repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: *repo,
	}
}

func (s *Service) GenerateLink(longUrl string) (interface{}, error) {

	urlHashBytes := sha256Generator(longUrl)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	if err != nil {
		return nil, err
	}
	err = s.repo.InsertLink(finalString[:8], longUrl)
	if err != nil {
		return nil, err
	}
	return finalString[:8], nil
}

func (s *Service) GetLongUrl(id string) (interface{}, error) {

	r, err := s.repo.GetLink(id)
	if err != nil {
		return nil, err
	}
	return r, nil
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
