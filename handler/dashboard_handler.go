package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func DashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		// c, err := r.Cookie("token")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		// If the cookie is not set, return an unauthorized status
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		return
		// 	}
		// 	// For any other type of error, return a bad request status
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// // Get the JWT string from the cookie
		// tknStr := c.Value
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tknStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Finally, return the welcome message to the user, along with their
		// username given in the token
		var s string = strconv.FormatUint(uint64(claims.UserId), 10)
		w.Write([]byte(fmt.Sprintf("Welcome %s!", s)))
	}
}
