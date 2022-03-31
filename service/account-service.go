package service

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/validators"
	"golang.org/x/crypto/bcrypt"
)

var AccountService *accountService

type accountService struct {
	*base
}

func (s *accountService) Initialize() {
	AccountService = &accountService{
		base: Base,
	}
}

func (s *accountService) Register(user domain.User) (domain.User, int, string) {
	// fullname valiation
	if strings.TrimSpace(user.Fullname) == "" {
		return user, http.StatusBadRequest, "Name cannot be white spaces"
	}
	// email validation
	if !validators.IsValidEmail(user.Email) {
		return user, http.StatusBadRequest, "Email is not a valid email"
	}
	// password validation
	//TODO:- currently the password is hard-coded, should not be in production
	user.Password = "123456Aa"
	// phone validation

	// encrypting password with bcrypt
	rand.Seed(time.Now().UnixNano())
	user.Cost = rand.Intn(20) + 1
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), user.Cost)
	user.Password = string(bytes)

	if err := s.repo.Create(user); err != nil {
		return user, http.StatusBadRequest, err.Error()
	}
	return user, http.StatusCreated, "New account has been created"
}
