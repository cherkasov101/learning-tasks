package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"io/ioutil"
	user "module30/pkg/user"
	"net/http"
	"os"
	"strconv"
)

var (
	countId  = 1
	fileName = "../db/db.json"
)

type Service struct {
	Users map[int]*user.User
}

func (s *Service) SaveDB() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(s.Users)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadDB() error {
	db, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(db, s.Users); err != nil {
		fmt.Println("json")
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

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

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
	id, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	if s.Users[id] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	for _, i := range s.Users[id].Friends {
		w.Write([]byte(strconv.Itoa(i)))
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Service) UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	if s.Users[id] == nil {
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

	s.Users[id].Age = update.NewAge
	w.Write([]byte("Возраст пользователя " + s.Users[id].Name + " обновлён: " + s.Users[id].Age))
	w.WriteHeader(http.StatusOK)
}
