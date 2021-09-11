package access_token

import (
	"strings"
	"time"

	"bookstore_oauth-api/src/domain/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")

	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid user id")

	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")

	}
	return nil

}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid gran_type parameter")
	}
	//TODO:Validate parameters for each request
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {

	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
