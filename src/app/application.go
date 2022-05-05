package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/http"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/repository/db"
	"github.com/mustafakocatepe/go-bookstore-oauth-api/src/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
