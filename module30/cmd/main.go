package main

import (
	"github.com/go-chi/chi/v5"
	service "module30/pkg/service"
	user "module30/pkg/user"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	srv := service.Service{make(map[int]*user.User)}
	r.Post("/create", srv.Create)
	r.Post("/make_friends", srv.MakeFriends)
	r.Delete("/delete_user/{id}", srv.Delete)
	r.Get("/friends/{id}", srv.GetUserFriends)
	r.Put("/new_age/{id}", srv.UpdateUserAge)

	http.ListenAndServe("localhost:8080", r)
}
