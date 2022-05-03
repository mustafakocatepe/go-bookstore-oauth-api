package db

import (
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/clients/cassandra"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}
