package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"interimProject/pkg/city"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	CountId string
	Cities  map[string]*city.City
}

func MakeService(db [][]string) *Service {
	cities := make(map[string]*city.City)
	intId := 0
	for _, row := range db {
		id := row[0]
		city := city.City{
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
		}
		cities[id] = &city
		newId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}
		if newId > intId {
			intId = newId
		}
	}
	intId++

	service := Service{
		strconv.Itoa(intId),
		cities,
	}

	return &service
}

func (s *Service) GetDBAsArray() [][]string {
	var db [][]string
	for i, city := range s.Cities {
		var cityArray []string
		cityArray = append(cityArray, i, city.Name, city.Region, city.District, city.Population, city.Foundation)
		db = append(db, cityArray)
	}
	return db
}

func (s *Service) GetCityInfo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if s.Cities[id] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	city := s.Cities[id]
	answer := "Название города: " + city.Name +
		"\nРегион: " + city.Region +
		"\nОкруг: " + city.District +
		"\nЧисленность населения: " + city.Population +
		"\nГод основания: " + city.Foundation

	w.Write([]byte(answer))
	w.WriteHeader(http.StatusOK)
}

func (s *Service) AddCity(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var city city.City
	if err := json.Unmarshal(content, &city); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Cities[s.CountId] = &city

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("id нового города: " + s.CountId + "\n"))

	count, err := strconv.Atoi(s.CountId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	count++
	s.CountId = strconv.Itoa(count)
}

func (s *Service) DeleteCity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if s.Cities[id] == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Неверный id"))
		return
	}

	w.Write([]byte("Город " + s.Cities[id].Name + " удалён"))
	delete(s.Cities, id)
	w.WriteHeader(http.StatusOK)
}

type UpdatePopulation struct {
	NewPopulation string `json:"new_population"`
}

func (s *Service) UpdatePopulation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if s.Cities[id] == nil {
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

	var update UpdatePopulation
	if err := json.Unmarshal(content, &update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.Cities[id].Population = update.NewPopulation
	w.Write([]byte("Информации о численности населения города " + s.Cities[id].Name + " обновлена: " + s.Cities[id].Population))
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetCitiesInRegion(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	cities := ""
	for _, city := range s.Cities {
		if city.Region == region {
			if cities != "" {
				cities += "\n"
			}
			cities += city.Name
		}
	}
	if cities == "" {
		cities = "Города в регионе " + region + " не найдены"
	}

	w.Write([]byte(cities))
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetCitiesInDistrict(w http.ResponseWriter, r *http.Request) {
	district := chi.URLParam(r, "district")
	cities := ""
	for _, city := range s.Cities {
		if city.District == district {
			if cities != "" {
				cities += "\n"
			}
			cities += city.Name
		}
	}
	if cities == "" {
		cities = "Города в округе " + district + " не найдены"
	}

	w.Write([]byte(cities))
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetCitiesByPopulation(w http.ResponseWriter, r *http.Request) {
	populationRange := chi.URLParam(r, "range")
	population := strings.Split(populationRange, "-")
	if len(population) != 2 {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cities := ""
	min, err := strconv.Atoi(population[0])
	if err != nil {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	max, err := strconv.Atoi(population[1])
	if err != nil {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, city := range s.Cities {
		popul, err := strconv.Atoi(city.Population)
		if err != nil {
			w.Write([]byte("db err"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if popul >= min && popul <= max {
			if cities != "" {
				cities += "\n"
			}
			cities += city.Name
		}
	}
	if cities == "" {
		cities = "Города с населением в диапазоне между " + population[0] + " и " + population[1] + " не найдены"
	}

	w.Write([]byte(cities))
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetCitiesByYears(w http.ResponseWriter, r *http.Request) {
	yearsRange := chi.URLParam(r, "range")
	years := strings.Split(yearsRange, "-")
	if len(years) != 2 {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cities := ""
	min, err := strconv.Atoi(years[0])
	if err != nil {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	max, err := strconv.Atoi(years[1])
	if err != nil {
		w.Write([]byte("Некорректный диапазон"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, city := range s.Cities {
		year, err := strconv.Atoi(city.Foundation)
		if err != nil {
			w.Write([]byte("db err"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if year >= min && year <= max {
			if cities != "" {
				cities += "\n"
			}
			cities += city.Name
		}
	}
	if cities == "" {
		cities = "Города с датой основания в диапазоне между " + years[0] + " и " + years[1] + " не найдены"
	}

	w.Write([]byte(cities))
	w.WriteHeader(http.StatusOK)
}
