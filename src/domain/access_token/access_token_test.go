package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T){
	assert.EqualValues(t, 24,expirationTime,"expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T){
	at := GetNewAccessToken()
	assert.False(t,at.IsExpired(),"brand new access token should be expired")
	assert.EqualValues(t,"",at.AccessToken,"new access token should not have defined access token id")
	assert.True(t,at.UserId==0,"new access token should not have an associated user id")

	//if at.AccessToken !=""{
	//	t.Error("new access token should not have defined access token id")
	//
	//}

	//if at.UserId!=0{
	//	t.Error("new access token should not have an associated user id")
	//}
}


func TestAccessTokenIsExpired(t *testing.T){
	at:=AccessToken{}
	assert.True(t,at.IsExpired()," access token should be expired by default")

	//if!at.IsExpired(){
	//	t.Error(" access token should be expired by default")
	//
	//}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t,at.IsExpired(),"access token expiring three hours form now should NOT be expired")

	//if at.IsExpired(){
	//	t.Error("access token expiring three hours form now should NOT be expired")
	//}

}