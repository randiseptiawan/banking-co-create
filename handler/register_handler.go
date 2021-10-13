package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/mysql"
)

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user mysql.User
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
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
			responder.NewHttpResponse(r, w, http.StatusBadRequest, err, nil)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusCreated, user, nil)
	}
}
