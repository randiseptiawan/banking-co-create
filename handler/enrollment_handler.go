package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/mysql"
)

func ReadAllUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := mysql.GetAllUser()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		for i := 0; i <= len(user)-1; i++ {
			user[i].Password = ""
		}
		responder.NewHttpResponse(r, w, http.StatusOK, user, nil)
	}
}

func UpdateEnrollmentStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tknStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
				return
			}
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		if !tkn.Valid {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}
		if claims.RoleStatus != "Admin" {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}
		var user mysql.User
		args := mux.Vars(r)
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &user)

		i, _ := strconv.ParseUint(args["id"], 10, 64)

		err = mysql.UpdateUser(uint(i), user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		userUpdated, err := mysql.GetUserById(uint(i))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		userUpdated.Password = ""
		responder.NewHttpResponse(r, w, http.StatusOK, userUpdated, nil)
	}
}
