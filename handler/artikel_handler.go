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

func CreateArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var artikel mysql.Artikel
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &artikel)

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
		artikel.UserId = claims.UserId
		err = mysql.CreateArtikel(&artikel)

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikelUser, err := mysql.GetUserById(artikel.UserId)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikel.UserName = artikelUser.NamaLengkap
		artikel.UserEmail = artikelUser.Email
		responder.NewHttpResponse(r, w, http.StatusCreated, artikel, nil)
	}
}

func ReadArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		i, _ := strconv.ParseUint(args["id"], 10, 64)

		artikel, err := mysql.ReadArtikel(i)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		artikelUser, err := mysql.GetUserById(artikel.UserId)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikel.UserName = artikelUser.NamaLengkap
		artikel.UserEmail = artikelUser.Email

		responder.NewHttpResponse(r, w, http.StatusOK, artikel, nil)
	}
}

func ReadAllArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		artikel, err := mysql.ReadAllArtikel()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		for i := 0; i <= len(artikel)-1; i++ {
			artikelUser, err := mysql.GetUserById(artikel[i].UserId)
			if err != nil {
				responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
				return
			}
			artikel[i].UserName = artikelUser.NamaLengkap
			artikel[i].UserEmail = artikelUser.Email
		}

		responder.NewHttpResponse(r, w, http.StatusOK, artikel, nil)
	}
}

func DeleteArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		i, _ := strconv.ParseUint(args["id"], 10, 64)

		err := mysql.DeleteArtikel(i)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}

func UpdateArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var artikel mysql.Artikel
		args := mux.Vars(r)
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &artikel)

		i, _ := strconv.ParseUint(args["id"], 10, 64)

		err = mysql.UpdateArtikel(uint(i), artikel)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikelUpdated, err := mysql.ReadArtikel(i)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikelUser, err := mysql.GetUserById(artikelUpdated.UserId)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		artikelUpdated.UserName = artikelUser.NamaLengkap
		artikelUpdated.UserEmail = artikelUser.Email
		responder.NewHttpResponse(r, w, http.StatusOK, artikelUpdated, nil)
	}
}
