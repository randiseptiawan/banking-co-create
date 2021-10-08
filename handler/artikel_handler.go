package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		json.Unmarshal(payloads, &artikel)

		err = mysql.CreateArtikel(artikel)

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
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
		responder.NewHttpResponse(r, w, http.StatusOK, artikel, nil)
	}
}

func DeleteArtikelHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		i, _ := strconv.ParseUint(args["id"], 10, 64)

		err := mysql.DeleteArtikel(i)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}
