package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	user "module31/pkg/user"
	"net/http"
	"os"
	"strconv"
)

var FileName = "../db/db.json"

type Service struct {
	CountId string                `json:"countId"`
	Users   map[string]*user.User `json:"users"`
}

func (s *Service) SaveDB() error {
	file, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

type MakingFriends struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

type UpdateAge struct {
	NewAge string `json:"new age"`
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

	if s.CountId == "" {
		s.CountId = "1"
	}

	var user user.User
	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Users[s.CountId] = &user

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("id нового пользователя: " + s.CountId + "\n"))

	count, err := strconv.Atoi(s.CountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	count++
	s.CountId = strconv.Itoa(count)

	if err = s.SaveDB(); err != nil {
		w.Write([]byte("x3"))
	}
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

	sourceId := makingFriend.SourceId
	targetId := makingFriend.TargetId

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

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id := userId

	if s.Users[id] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	w.Write([]byte("Пользователь " + s.Users[id].Name + " удалён"))
	s.Users[id] = nil
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetUserFriends(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id := userId

	if s.Users[id] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	for _, i := range s.Users[id].Friends {
		w.Write([]byte(i))
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Service) UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if s.Users[userId] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var update UpdateAge
	if err := json.Unmarshal(content, &update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Users[userId].Age = update.NewAge
	w.Write([]byte("Возраст пользователя " + s.Users[userId].Name + " обновлён: " + s.Users[userId].Age))
	w.WriteHeader(http.StatusOK)
}
