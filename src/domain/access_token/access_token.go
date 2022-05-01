package access_token

import "time"

const (
	expirationtime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` //Web Frontend veya Mobil cihazlara destek verebilirim ve bunlarda ki Expire zamanlarım Token için farklı olabilir. Bu yüzden bu bilgiyi tutuyorum.
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationtime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {

	now := time.Now().UTC()
	expirationtime := time.Unix(at.Expires, 0)
	return now.After(expirationtime)
}
