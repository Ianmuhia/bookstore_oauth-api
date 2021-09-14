package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	atDomain "bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/domain/utils/errors"
	"bookstore_oauth-api/src/services/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
	// var at atDomain.AccessToken
	// if err := c.ShouldBindJSON(&at); err != nil {
	// 	restErr := errors.NewBadRequestError("invalid json body")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }
	// if err := handler.service.Create(at); err != nil {
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	// c.JSON(http.StatusCreated, at)
}
