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

type MakingFriends struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
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
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var makingFriend MakingFriends
	if err := json.Unmarshal(content, &makingFriend); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sourceId, err := strconv.Atoi(makingFriend.SourceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	targetId, err := strconv.Atoi(makingFriend.TargetId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	if s.Users[sourceId] == nil || s.Users[targetId] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не существует"))
		return
	}

	s.Users[sourceId].Friends = append(s.Users[sourceId].Friends, targetId)
	s.Users[targetId].Friends = append(s.Users[targetId].Friends, sourceId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s.Users[sourceId].Name + " и " + s.Users[targetId].Name + " теперь друзья"))
}
