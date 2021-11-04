package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/external/mysql"
	"golang.org/x/crypto/bcrypt"
)

type M map[string]interface{}

var jwtKey = []byte(config.Init().JWTSecret)

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

//
type Claims struct {
	UserId      uint   `json:"UserId"`
	NamaLengkap string `json:"NamaLengkap"`
	Email       string `json:"Email"`
	RoleStatus  string `json:"RoleStatus"`
	jwt.StandardClaims
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		user, err := mysql.GetUser(creds.Email)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}

		expirationTime := time.Now().Add(60 * time.Minute)
		claims := &Claims{
			UserId:      user.Model.ID,
			NamaLengkap: user.NamaLengkap,
			Email:       user.Email,
			RoleStatus:  user.RoleStatus,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString(jwtKey)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
			return
		}
		tokenString, _ := json.Marshal(M{"token": signedToken})
		w.Write([]byte(tokenString))
		// responder.NewHttpResponse(r, w, http.StatusCreated, []byte(tokenString), nil)
	}
}

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user mysql.User
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &user)
		var passwordBytes = []byte(user.Password)

		hashedPassword, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

		user.Password = string(hashedPassword)
		user.EnrollmentStatus = "Waiting for Approval"
		user.RoleStatus = "Member"

		err = mysql.CreateUser(&user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		user.Password = ""
		responder.NewHttpResponse(r, w, http.StatusCreated, user, nil)
	}
}

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := []byte("Welcome")
		responder.NewHttpResponse(r, w, http.StatusOK, res, nil)
	}
}

func RegisterAdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user mysql.User
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &user)
		var passwordBytes = []byte(user.Password)

		hashedPassword, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

		user.Password = string(hashedPassword)
		user.EnrollmentStatus = "Waiting for Approval"
		user.RoleStatus = "Admin"

		err = mysql.CreateUser(&user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		user.Password = ""
		responder.NewHttpResponse(r, w, http.StatusCreated, user, nil)
	}
}
