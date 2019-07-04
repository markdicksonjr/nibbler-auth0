package main

import (
	"github.com/markdicksonjr/nibbler"
	"github.com/markdicksonjr/nibbler-auth0/user"
	"log"
	"net/http"
)

type SampleExtension struct {
	nibbler.NoOpExtension
	Auth0Extension *user.Extension
}

func (s *SampleExtension) AddRoutes(app *nibbler.Application) error {
	app.GetRouter().HandleFunc("/test", s.Auth0Extension.EnforceLoggedIn(s.ProtectedRoute)).Methods("GET")
	return nil
}

func (s *SampleExtension) ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	u, err := s.Auth0Extension.SessionExtension.GetCaller(r)

	if err != nil {
		nibbler.Write500Json(w, err.Error())
		return
	}

	if u == nil {
		nibbler.Write404Json(w)
		return
	}

	log.Println(u)

	nibbler.Write200Json(w, `{"result": "authorized"}`)
}
