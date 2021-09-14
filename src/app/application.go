package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"bookstore_oauth-api/src/clients/cassandra"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"
	"bookstore_oauth-api/src/repository/rest"
	"bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()

	session.Close()

	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
