package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/external/mysql"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(config.Init().JWTSecret)

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

//
type Claims struct {
	NamaLengkap string `json:"namaLengkap"`
	Email       string `json:"email"`
	RoleStatus  string `json:"roleStatus`
	jwt.StandardClaims
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// Get the JSON body and decode into credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the expected password from our in memory map
		// expectedPassword, ok := users[creds.Username]
		user, _ := mysql.GetUser(creds.Email)

		// // If a password exists for the given user
		// // AND, if it is the same as the password we received, the we can move ahead
		// // if NOT, then we return an "Unauthorized" status
		// if !ok || user.Password != creds.Password {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			w.WriteHeader(http.StatusUnauthorized)
		}

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(60 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			NamaLengkap: user.NamaLengkap,
			Email:       user.Email,
			RoleStatus:  user.RoleStatus,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}
}
