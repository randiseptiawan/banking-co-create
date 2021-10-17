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
	UserId     uint   `json:"userId"`
	Email      string `json:"email"`
	RoleStatus string `json:"roleStatus`
	jwt.StandardClaims
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// Get the JSON body and decode into credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		// Get the expected password from our in memory map
		// expectedPassword, ok := users[creds.Username]
		user, err := mysql.GetUser(creds.Email)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}
		// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.MinCost)
		// fmt.Println(string(hashedPassword), user.Password)
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(60 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			UserId:     user.Model.ID,
			Email:      user.Email,
			RoleStatus: user.RoleStatus,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		signedToken, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
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

		err = mysql.Register(&user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusCreated, user, nil)
	}
}
