package rest

import (
	"encoding/json"
	resty "github.com/go-resty/resty/v2"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/domain/users"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/utils/errors"
	"time"
)

var (
	usersRestClient = resty.New().SetBaseURL("http://localhost:8082").SetTimeout(100 * time.Millisecond)
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type restUsersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response, err := usersRestClient.R().
		SetBody(request).
		Post("/users/login")

	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode() > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil

}
