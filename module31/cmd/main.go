package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	service "module31/pkg/service"
	"net/http"
)

var localHost = "localhost:"

func main() {
	var lh string
	flag.StringVar(&lh, "lh", "8080", "set localhost")
	flag.Parse()

	localHost += lh

	r := chi.NewRouter()

	var srv service.Service
	db, err := ioutil.ReadFile(service.FileName)
	if err != nil {
		fmt.Println("can't read")
		return
	}
	if err = json.Unmarshal(db, &srv); err != nil {
		fmt.Println("can't read json")
		return
	}

	r.Post("/create", srv.Create)
	r.Post("/make_friends", srv.MakeFriends)
	r.Delete("/delete_user/{id}", srv.Delete)
	r.Get("/friends/{id}", srv.GetUserFriends)
	r.Put("/new_age/{id}", srv.UpdateUserAge)

	http.ListenAndServe(localHost, r)
}
