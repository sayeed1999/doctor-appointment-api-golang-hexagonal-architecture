package handlers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/service"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/helpers"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/pkg"
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

	if err := pkg.DecodeJSON(c.Request.Body, &user); err != nil {
		helpers.Response(c, 400, "", nil)
		return
	}

	user, code, text := service.AccountService.Register(user)
	helpers.Response(c, code, text, user)
}

func (h *accountHandler) Login(c *gin.Context) {
	// var user domain.User
	var loginData map[string]string

	if err := pkg.DecodeJSON(c.Request.Body, &loginData); err != nil {
		helpers.Response(c, 400, "", nil)
		return
	}

	email := loginData["email"]
	password := loginData["password"]

	token, code, text := service.AccountService.Login(email, password)
	// fmt.Println("token: ", token)
	c.SetCookie("jwttoken", token, 60*60*24, "/", "localhost", false, true)

	helpers.Response(c, code, text, nil)
}

func (h *accountHandler) GetAuthenticatedUser(c *gin.Context) {
	cookie, err := c.Cookie("jwttoken")
	if err != nil {
		helpers.Response(c, http.StatusUnauthorized, "", nil)
		// helpers.Response(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	claims := jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.conf.Jwt.SecretKey), nil
	})
	if err != nil {
		helpers.Response(c, http.StatusUnauthorized, "", nil)
		return
	}

	// token := tokenWithClaims.Raw; fmt.Println(token)
	// if claims.ExpiresAt <= time.Now().Unix() {
	// 	helpers.Response(c, http.StatusUnauthorized, "Login state has been expired!", nil)
	// 	return
	// }

	userid, _ := strconv.ParseInt(claims.Issuer, 10, 32)
	user, code, text := service.AccountService.GetUserById(int(userid))

	helpers.Response(c, code, text, user)
	return
}
