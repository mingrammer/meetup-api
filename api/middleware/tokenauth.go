package middleware

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
)

type TokenAuthMiddleware struct {
	Realm         string
	Authenticator func(token string) bool
	Authorizator  func(token string, request *rest.Request) bool
}

func (mw *TokenAuthMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	if mw.Realm == "" {
		log.Fatal("Realm is required")
	}
	if mw.Authenticator == nil {
		log.Fatal("Authenticator is required")
	}
	if mw.Authorizator == nil {
		mw.Authorizator = func(token string, request *rest.Request) bool {
			return true
		}
	}
	return func(writer rest.ResponseWriter, request *rest.Request) {
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			mw.unauthorized(writer)
			return
		}
		providedToken, err := mw.decodeBasicAuthHeader(authHeader)
		if err != nil {
			rest.Error(writer, "Invalid authentication", http.StatusBadRequest)
			return
		}
		if !mw.Authenticator(providedToken) {
			mw.unauthorized(writer)
			return
		}
		if !mw.Authorizator(providedToken, request) {
			mw.unauthorized(writer)
			return
		}
		request.Env["VALID_USER_TOKEN"] = providedToken
		handler(writer, request)
	}
}

func (mw *TokenAuthMiddleware) unauthorized(writer rest.ResponseWriter) {
	writer.Header().Set("WWW-Authenticate", "Basic realm="+mw.Realm)
	rest.Error(writer, "Not Authorized", http.StatusUnauthorized)
}

func (mw *TokenAuthMiddleware) decodeBasicAuthHeader(header string) (tokwn string, err error) {
	parts := strings.SplitN(header, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Token") {
		return "", errors.New("Invalid Authorization header")
	}
	_, err = base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", errors.New("Token encoding not valid")
	}
	return string(parts[1]), nil
}
