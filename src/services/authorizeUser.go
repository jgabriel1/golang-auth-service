package services

import (
	"errors"
	"golang-auth-service/src/repo"
	"golang-auth-service/src/security"
	"net/http"
	"regexp"
	"strings"
)

type AuthorizeUser struct {
	usersRepo *repo.UsersRepository
}

type AuthCredentials struct {
	UserID string
}

func NewAuthorizeUser(usersRepo *repo.UsersRepository) *AuthorizeUser {
	this := AuthorizeUser{
		usersRepo: usersRepo,
	}

	return &this
}

func (this *AuthorizeUser) parseTokenValue(authHeadersValue string) (string, error) {
	jwtRegExp, _ := regexp.Compile(`^Bearer\s[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`)

	if match := jwtRegExp.Match([]byte(authHeadersValue)); !match {
		return "", errors.New("Authorization header has the wrong format.")
	}

	// "Bearer <token>"
	tokenValue := strings.Fields(authHeadersValue)[1]

	return tokenValue, nil
}

func (this *AuthorizeUser) parseHeaders(headers *http.Header) (string, error) {
	authHeader := headers.Get("authorization")
	if authHeader == "" {
		return "", errors.New("Missing \"Authorization\" headers.")
	}

	return this.parseTokenValue(authHeader)
}

func (this *AuthorizeUser) Execute(requestHeaders *http.Header) (*AuthCredentials, error) {
	tokenValue, err := this.parseHeaders(requestHeaders)
	if err != nil {
		return nil, err
	}

	jwtPayload, err := security.ValidateJWT(tokenValue)
	if err != nil {
		return nil, err
	}

	credentials := AuthCredentials{
		UserID: jwtPayload.UserID,
	}

	return &credentials, nil
}
