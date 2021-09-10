package rest

import "bookstore_oauth-api/src/domain/utils/errors"

type RestUserRepository interface {
	LoginUser(string, string) (*User, errors.RestErr)
}

func NewRepository() RestUserRepository {
	return &RestUserRepository{}
}
