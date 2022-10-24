package service

import (
	"encoding/json"
	"io"
	user "module30/pkg/user"
	"net/http"
	"strconv"
)

var countId = 1

type Service struct {
	Users map[int]*user.User
}

// Create - function for creating new user.
func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var user user.User
	user.Friends = make([]int, 5)
	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Users[countId] = &user

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("id нового пользователя: " + strconv.Itoa(countId) + "\n"))

	countId++
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {

}
