package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	setCheckRouter(r, dependencies.CheckService)
	setArtikelRouter(r)
	setProjectRouter(r)
	setLoginRouter(r)
	setDashboardRouter(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func setCheckRouter(router *mux.Router, checkService service.CheckService) {
	router.Methods(http.MethodGet).Path("/check/redis").Handler(handler.CheckRedis(checkService))
	router.Methods(http.MethodGet).Path("/check/mysql").Handler(handler.CheckMysql(checkService))
	router.Methods(http.MethodGet).Path("/check/minio").Handler(handler.CheckMinio(checkService))
}

func setArtikelRouter(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/artikel/{id}").Handler(handler.ReadArtikelHandler())
	router.Methods(http.MethodGet).Path("/artikel").Handler(handler.ReadAllArtikelHandler())
	router.Methods(http.MethodPost).Path("/artikel").Handler(handler.CreateArtikelHandler())
	router.Methods(http.MethodDelete).Path("/artikel/{id}").Handler(handler.DeleteArtikelHandler())
	router.Methods(http.MethodPut).Path("/artikel/{id}").Handler(handler.UpdateArtikelHandler())
}

func setProjectRouter(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/project/{id}").Handler(handler.ReadProjectHandler())
	router.Methods(http.MethodGet).Path("/project").Handler(handler.ReadAllProjectHandler())
	router.Methods(http.MethodPost).Path("/project").Handler(handler.CreateProjectHandler())
	router.Methods(http.MethodDelete).Path("/project/{id}").Handler(handler.DeleteProjectHandler())
	router.Methods(http.MethodPost).Path("/project/invite/{id}").Handler(handler.InviteUserHandler())
	router.Methods(http.MethodPut).Path("/project/{id}").Handler(handler.UpdateProjectHandler())
}

func setLoginRouter(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/login").Handler(handler.LoginHandler())
	router.Methods(http.MethodPost).Path("/register").Handler(handler.RegisterHandler())
}

func setDashboardRouter(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/").Handler(handler.HomeHandler())
}
