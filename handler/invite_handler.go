package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/mysql"
)

func InviteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inv Invite
		// Get the JSON body and decode into credentials
		err := json.NewDecoder(r.Body).Decode(&inv)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		// Get the expected password from our in memory map
		// expectedPassword, ok := users[creds.Username]
		user, err := mysql.GetUser(inv.Email)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		args := mux.Vars(r)
		IdProject, _ := strconv.ParseUint(args["id"], 10, 64)

		var invited mysql.Invited
		invited.ProjectId = uint(IdProject)
		invited.InvitedUserId = user.Model.ID
		err = mysql.InviteUser(&invited)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusCreated, invited, nil)
	}
}

func AcceptInvitedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var collaborator mysql.Collaborator
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &collaborator)
		fmt.Println(collaborator)
		err = mysql.DeleteInvitedUser(uint64(collaborator.CollaboratorUserId), uint64(collaborator.ProjectId))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		err = mysql.CreateCollaborator(&collaborator)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, nil, nil)
	}
}

func IgnoreInvitedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var collaborator mysql.Collaborator
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &collaborator)

		err = mysql.DeleteInvitedUser(uint64(collaborator.CollaboratorUserId), uint64(collaborator.ProjectId))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, nil, nil)
	}
}

func ReadProjectInvitedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		project, err := mysql.ProjectInvited(claims.UserId)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		projectAdmin, err := mysql.GetUserById(project.ProjectAdminId)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		project.ProjectAdminName = projectAdmin.NamaLengkap
		project.ProjectAdminEmail = projectAdmin.Email

		responder.NewHttpResponse(r, w, http.StatusOK, project, nil)
	}
}
