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
	setRegisterRouter(r)
	setArtikelRouter(r)

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

func setRegisterRouter(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/register").Handler(handler.Register())
}

func setArtikelRouter(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/artikel/detail/{id}").Handler(handler.ReadArtikelHandler())
	router.Methods(http.MethodGet).Path("/list_artikel").Handler(handler.ReadAllArtikelHandler())
	router.Methods(http.MethodPost).Path("/artikel/create").Handler(handler.CreateArtikelHandler())
	router.Methods(http.MethodDelete).Path("/artikel/delete/{id}").Handler(handler.DeleteArtikelHandler())
}