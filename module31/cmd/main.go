package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	service "module31/pkg/service"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	var srv service.Service
	db, err := ioutil.ReadFile(service.FileName)
	if err != nil {
		fmt.Println("can't read")
	}
	if err = json.Unmarshal(db, &srv); err != nil {
		fmt.Println("json")
	}
	fmt.Println(srv.CountId)
	r.Post("/create", srv.Create)
	r.Post("/make_friends", srv.MakeFriends)
	r.Delete("/delete_user/{id}", srv.Delete)
	r.Get("/friends/{id}", srv.GetUserFriends)
	r.Put("/new_age/{id}", srv.UpdateUserAge)

	http.ListenAndServe("localhost:8080", r)
}
