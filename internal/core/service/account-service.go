package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	cons "github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/constants"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/repository"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	*base
	repo *repository.AccountRepository
}

func InitializeAccountService(b *base, r *repository.AccountRepository) *accountService {
	return &accountService{
		base: b,
		repo: r,
	}
}

func (s *accountService) Register(user domain.User) (domain.User, int, string) {
	// fullname valiation
	if strings.TrimSpace(user.Fullname) == "" {
		return user, http.StatusBadRequest, "Name cannot be white spaces"
	}
	// email validation
	if !validators.IsValidEmail(user.Email) {
		return user, http.StatusBadRequest, cons.ApplicationMessage.InvalidEmail
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

func (s *accountService) Login(email string, password string) (string, int, string) {
	var user domain.User

	// find if the email exists
	if err := s.repo.FindFirst(&user, "email = ?", email); err != nil {
		return "", http.StatusInternalServerError, ""
	}
	if user.ID == 0 {
		return "", http.StatusNotFound, cons.ApplicationMessage.EmailNotRegistered
	}

	// hashedPasswordInBytes, _ := bcrypt.GenerateFromPassword([]byte(password), user.Cost)
	// hashedPassword := string(hashedPasswordInBytes)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", http.StatusBadRequest, cons.ApplicationMessage.WrongPassword
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    strconv.Itoa(int(user.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(s.conf.Jwt.SecretKey))

	return signedToken, http.StatusAccepted, ""
}

func (s *accountService) GetUserById(id int) (domain.User, int, string) {
	var user domain.User

	_ = s.repo.FindById(&user, id)

	if user.ID == 0 {
		return user, http.StatusNotFound, fmt.Sprintf(cons.ApplicationMessage.ItemNotFound, "User")
	}
	return user, http.StatusFound, ""
}
