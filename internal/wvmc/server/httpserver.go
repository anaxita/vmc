package server

import (
	"net/http"
	"os"

	"github.com/anaxita/logit"
	"github.com/anaxita/wvmc/internal/wvmc/store"
	"github.com/gorilla/mux"
)

// Server - структура http сервера
type Server struct {
	store  *store.Store
	router *mux.Router
}

// New - создает новый сервер
func New(storage *store.Store) *Server {
	return &Server{store: storage,
		router: mux.NewRouter()}
}

func (s *Server) configureRouter() {
	r := s.router
	r.Use(s.Cors)
	r.Handle("/refresh", s.RefreshToken()).Methods("POST", "OPTIONS")
	r.Handle("/signin", s.SignIn()).Methods("POST", "OPTIONS")

	users := r.NewRoute().Subrouter()
	users.Use(s.Auth, s.CheckIsAdmin)
	users.Handle("/users", s.GetUsers()).Methods("OPTIONS", "GET")
	users.Handle("/users", s.CreateUser()).Methods("POST", "OPTIONS")
	users.Handle("/users", s.EditUser()).Methods("OPTIONS", "PATCH")
	users.Handle("/users", s.DeleteUser()).Methods("OPTIONS", "DELETE")
	users.Handle("/users/servers", s.AddServerToUser()).Methods("OPTIONS", "POST")

	serversShow := r.NewRoute().Subrouter()
	serversShow.Use(s.Auth)
	serversShow.Handle("/servers", s.GetServers()).Methods("OPTIONS", "GET")

	servers := r.NewRoute().Subrouter()
	servers.Use(s.Auth, s.CheckIsAdmin)
	servers.Handle("/servers", s.CreateUser()).Methods("POST", "OPTIONS")
	servers.Handle("/servers", s.EditUser()).Methods("OPTIONS", "PATCH")
	servers.Handle("/servers", s.DeleteUser()).Methods("OPTIONS", "DELETE")
}

// Start - запускает сервер
func (s *Server) Start() error {
	s.configureRouter()
	logit.Info("Server stared at", os.Getenv("PORT"))
	return http.ListenAndServe(os.Getenv("PORT"), s.router)
}
