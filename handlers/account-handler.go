package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/helpers"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/service"
)

var AccountHandler *accountHandler

type accountHandler struct {
	*base
}

func (h *accountHandler) Initialize() {
	AccountHandler = &accountHandler{
		base: Base,
	}
}

func (h *accountHandler) Register(c *gin.Context) {
	//showing that map[string]interface{} can extract any post form in golang
	// var data map[string]interface{}
	var user domain.User

	if err := helpers.DecodeJSON(c.Request.Body, &user); err != nil {
		helpers.Response(c, 400, "", nil)
		return
	}

	user, code, text := service.AccountService.Register(user)
	helpers.Response(c, code, text, user)
}
