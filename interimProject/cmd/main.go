package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-chi/chi/v5"
	service "interimProject/pkg/service"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	fileName  = "../db/cities.csv"
	db        [][]string
	localHost = "localhost:8080"
)

func main() {
	db = readDB()
	service := service.MakeService(db)

	go makeServer(service)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

	ForLoop:
		for {
			select {
			case <-done:
				db = service.GetDBAsArray()
				writeDB()
				break ForLoop
			}
		}
	}()

	wg.Wait()
}

func makeServer(service *service.Service) {
	r := chi.NewRouter()

	r.Get("/city_info/{id}", service.GetCityInfo)
	r.Post("/add_city", service.AddCity)
	r.Delete("/delete/{id}", service.DeleteCity)
	r.Put("/update_population/{id}", service.UpdatePopulation)
	r.Get("/cities_in_region/{region}", service.GetCitiesInRegion)
	r.Get("/cities_in_district/{district}", service.GetCitiesInDistrict)
	r.Get("/cities_by_population/{range}", service.GetCitiesByPopulation)
	r.Get("/cities_by_years/{range}", service.GetCitiesByYears)

	http.ListenAndServe(localHost, r)
}

func readDB() [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't read")
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	db, err := reader.ReadAll()
	if err != nil {
		fmt.Println("problem with reading")
		return nil
	}

	return db
}

func writeDB() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("can't create file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("can't write")
	}
}
