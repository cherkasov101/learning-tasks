package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter  = 0
	firstIH  = "http://localhost:8080"
	secondIH = "http://localhost:8081"
)

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/make_friends", makeFriends)
	http.HandleFunc("/delete_user/{id}", deleteUser)
	http.HandleFunc("/friends/{id}", getUserFriends)
	http.HandleFunc("/new_age/{id}", newAge)

	http.ListenAndServe(proxyAddr, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	var localHost string

	if counter == 0 {
		localHost = firstIH
		counter++
	} else {
		localHost = secondIH
		counter--
	}

	localHost += "/create"

	client := &http.Client{}

	req, err := http.NewRequest("POST", localHost, bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
}

func makeFriends(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	var localHost string

	if counter == 0 {
		localHost = firstIH
		counter++
	} else {
		localHost = secondIH
		counter--
	}

	localHost += "/make_friends"

	client := &http.Client{}

	req, err := http.NewRequest("POST", localHost, bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	var localHost string

	if counter == 0 {
		localHost = firstIH
		counter++
	} else {
		localHost = secondIH
		counter--
	}

	id := chi.URLParam(r, "id")

	localHost += "/delete_user/" + id

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", localHost, bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
}

func getUserFriends(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	var localHost string

	if counter == 0 {
		localHost = firstIH
		counter++
	} else {
		localHost = secondIH
		counter--
	}

	id := chi.URLParam(r, "id")
	localHost += "/friends/" + id

	client := &http.Client{}

	req, err := http.NewRequest("GET", localHost, bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
	w.Write(bodyBytes)
}

func newAge(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	var localHost string

	if counter == 0 {
		localHost = firstIH
		counter++
	} else {
		localHost = secondIH
		counter--
	}

	id := chi.URLParam(r, "id")
	localHost += "/new_age/" + id

	client := &http.Client{}

	req, err := http.NewRequest("PUT", localHost, bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
}
