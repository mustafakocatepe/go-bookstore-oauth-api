package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/domain/users"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/utils/errors"
	"io"
	"net/http"
)

var (
	Client  http.Client
	baseUrl = "http://localhost:8082"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type restUsersRepository struct {
	Client *http.Client
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{Client: &http.Client{}}
}

func (r *restUsersRepository) LoginUser(email string, password string) (user *users.User, error *errors.RestErr) {
	requestModel := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	url := fmt.Sprintf("%s/users/login", baseUrl)

	postBody, _ := json.Marshal(requestModel)
	responseBody := bytes.NewBuffer(postBody)

	request, _ := http.NewRequest(http.MethodPost, url, responseBody)
	request.Header.Add("Accept", "application/json")

	response, err := r.Client.Do(request)

	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")

	}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")

	}

	if response.StatusCode != 200 {
		var restErr errors.RestErr

		if err := json.Unmarshal(body, &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")

		}

		return nil, &restErr
	}

	return user, nil
}

/*
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

}*/
