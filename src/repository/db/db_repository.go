package db

import (
	"bookstore_oauth-api/src/clients/cassandra"
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/domain/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(*access_token.AccessToken) *errors.RestErr
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewInternalServerError("no access token found with the given id")

		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at *access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return nil
}
