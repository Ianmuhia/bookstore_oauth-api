package access_token

import "time"

const(
	expirationTime = 24
)

type AccessToken struct{
	AccessToken string `json:"access_token"`
	UserId int64  `json:"user_id"`
	ClientId int64 `json:"client_id"`
	ExpiresIn int64 `json:"expires"`

}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		ExpiresIn: time.Now().UTC().Add(expirationTime * time.Hour ).Unix(),
	}
}

func (at AccessToken) IsExpired()bool{
	return false
}