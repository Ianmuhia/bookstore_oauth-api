package rest

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	"bookstore_oauth-api/src/domain/users"
	"bookstore_oauth-api/src/domain/utils/errors"
)

var (
	client = resty.New()
	// usersRestClient = rest.RequestBuilder{
	// 	Headers:        nil,
	// 	Timeout:        100 * time.Millisecond,
	// 	ConnectTimeout: 0,
	// 	BaseURL:        "https://api.bookstore.com",
	// }
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type usersRepository struct {
}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (s *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	// client := resty.New()
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response, err := client.R().
		SetBody(request).
		//   SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		//   SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post("/users/login")
	if err != nil {
		panic(err)

	}

	fmt.Println(response.Body())
	// response := usersRestClient.Post("/users/login", request)
	// if response == nil || response.Response == nil {
	// 	return nil, errors.NewInternalServerError("invalid response when trying to login user")
	// }
	// if response.StatusCode > 299 {
	// 	var restErr errors.RestErr
	// 	err := json.Unmarshal(response.Bytes(), &restErr)
	// 	if err != nil {
	// 		return nil, errors.NewInternalServerError("invalid error interface when trying to login user")

	// 	}
	// 	return nil, &restErr
	// }
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
