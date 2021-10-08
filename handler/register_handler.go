package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"crypto/sha256"
	"encoding/base64"

	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/mysql"
)

func encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user mysql.User
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		json.Unmarshal(payloads, &user)
		user.Password = encryptPassword(user.Password)
		err = mysql.Register_Member(user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusCreated, user, nil)
	}
}
