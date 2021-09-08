package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T){
	at := GetNewAccessToken()
	if at.IsExpired(){
		t.Error("brand new access token should be expired")
	}

	if at.AccessToken !=""{
		t.Error("new access token should not have defined access token id")

	}
	if at.UserId!=0{
		t.Error("new access token should not have an associated user id")
	}
}


func TestAccessTokenIsExpired(t *testing.T){
	at:=AccessToken{}
	if!at.IsExpired(){
		t.Error(" access token should be expired by default")

	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired(){
		t.Error("access token expiring three hours form now should NOT be expired")
	}

}