package acces_token_service

import (
	"strings"

	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/domain/utils/errors"
	"bookstore_oauth-api/src/repository/db"
	"bookstore_oauth-api/src/repository/rest"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUserRepository
	dbRepo db.DbRepository
	// repository Repository
}

func NewService(usersRepo rest.RestUserRepository,dbRepo db.DbRepository ) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo: dbRepo,
		// repository: repo,

	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {

	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	//TODO:support both client credentials and password
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err != nil {
		return nil, err
	}
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil

}
func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
