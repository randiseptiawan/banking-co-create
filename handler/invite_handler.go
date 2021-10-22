package handler

import (
	"encoding/json"
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
		args := mux.Vars(r)
		IdProject, _ := strconv.ParseUint(args["id"], 10, 64)
		project, err := mysql.ReadProject(IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		if claims.UserId != project.ProjectAdminId {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}

		user, err := mysql.GetUser(inv.Email)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		var invited mysql.Invited
		invited.ProjectId = uint(IdProject)
		invited.InvitedUserId = user.Model.ID
		err = mysql.InviteUser(&invited)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectUpdated, err := mysql.ReadProject(IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectAdmin, err := mysql.GetProjectAdmin(uint(IdProject))
		projectUpdated.ProjectAdmin = projectAdmin

		invitedUser, err := mysql.GetInvitedUsername(uint(IdProject))
		projectUpdated.InvitedUser = invitedUser

		collaboratedUser, err := mysql.GetCollaboratedUsername(uint(IdProject))
		projectUpdated.Collaborator = collaboratedUser
		responder.NewHttpResponse(r, w, http.StatusCreated, projectUpdated, nil)
	}
}

func AcceptInvitedHandler() http.HandlerFunc {
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
		args := mux.Vars(r)
		IdProject, _ := strconv.ParseUint(args["id"], 10, 64)
		invitedUserProject, err := mysql.GetInvitedUser(claims.UserId, IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		if len(invitedUserProject) == 0 {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}

		err = mysql.DeleteInvitedUser(uint64(claims.UserId), uint64(IdProject))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		var collaborator mysql.Collaborator
		collaborator.ProjectId = uint(IdProject)
		collaborator.CollaboratorUserId = uint(claims.UserId)
		err = mysql.CreateCollaborator(&collaborator)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectUpdated, err := mysql.ReadProject(IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectAdmin, err := mysql.GetProjectAdmin(uint(IdProject))
		projectUpdated.ProjectAdmin = projectAdmin

		invitedUser, err := mysql.GetInvitedUsername(uint(IdProject))
		projectUpdated.InvitedUser = invitedUser

		collaboratedUser, err := mysql.GetCollaboratedUsername(uint(IdProject))
		projectUpdated.Collaborator = collaboratedUser
		responder.NewHttpResponse(r, w, http.StatusCreated, projectUpdated, nil)
	}
}

func IgnoreInvitedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
		args := mux.Vars(r)
		IdProject, _ := strconv.ParseUint(args["id"], 10, 64)
		invitedUserProject, err := mysql.GetInvitedUser(claims.UserId, IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		if len(invitedUserProject) == 0 {
			responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, err)
			return
		}
		err = mysql.DeleteInvitedUser(uint64(claims.UserId), uint64(IdProject))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectUpdated, err := mysql.ReadProject(IdProject)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		projectAdmin, err := mysql.GetProjectAdmin(uint(IdProject))
		projectUpdated.ProjectAdmin = projectAdmin

		invitedUser, err := mysql.GetInvitedUsername(uint(IdProject))
		projectUpdated.InvitedUser = invitedUser

		collaboratedUser, err := mysql.GetCollaboratedUsername(uint(IdProject))
		projectUpdated.Collaborator = collaboratedUser
		responder.NewHttpResponse(r, w, http.StatusCreated, projectUpdated, nil)
	}
}
