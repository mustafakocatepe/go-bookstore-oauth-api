package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	//Testlerimiz de kontrollerimizi aşağıda ki şekilde kontrol etmek yerine, assert package ile yapabiliriz.

	/*if expirationtime != 24 {
		t.Error()
	}*/

	assert.EqualValues(t, 24, expirationtime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	/*if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}*/
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	/*if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}*/
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")

	/*if at.UserId != 0 {
		t.Error("new access token should not have associated user id")
	}*/
	assert.Truef(t, at.UserId == 0, "new access token should not have associated user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	/*if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}*/
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	/*if at.IsExpired() {
		t.Error("access token expiring  three hours  from now should not be expired be expired")
	}*/
	assert.Falsef(t, at.IsExpired(), "access token expiring  three hours  from now should not be expired be expired")

}
