package accesstoken

import (
	"strings"
	"time"

	"github.com/ankitanwar/OAuth/utils/errors"
)

const (
	experationTime = 24
)

//AccessToken : fields of access token
type AccessToken struct {
	AccessToken string `json:"access_token" bson:"access_token"`
	UserID      int    `json:"user_id" bson:"user_id"`
	ClinetID    int    `json:"client_id" bson:"clinet_id"` //to determine whether the client is from mobile app or web
	Expires     int64  `json:"expires" bson:"expires"`
}

//TokenRequest : To Request The Acees Token
type TokenRequest struct  {
	Email          string `json:"email"`
	Password string `json:"password"`
}

//Validate : To validate the Access Token
func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewInternalServerError("Invalid Access Token")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequest("Invalid User ID")
	}
	if at.ClinetID <= 0 {
		return errors.NewBadRequest("Invalid Clinet ID")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequest("Invalid experation ID")
	}
	return nil
}

//GetNewAccess : To get the new access token
func GetNewAccess() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(experationTime * time.Hour).Unix(),
	}
}

//IsExpired : To Check whether the given access token is experied or not
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
