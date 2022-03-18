package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	gocron "github.com/go-co-op/gocron"
)

type Weather struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var PORT = ":8080"

func main() {
	fmt.Println("MAIN")
	http.HandleFunc("/", getWeather)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)

}

func getWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getWeather")

	s := gocron.NewScheduler(time.UTC)

	if r.Method == "GET" {

		//Buka json file
		jsonFile, err := os.Open("data.json")

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Successfully Opened data.json")

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var weather Weather

		json.Unmarshal(byteValue, &weather)
		fmt.Println("BEFORE")

		s.Every(5).Seconds().Do(func() {
			fmt.Println("MASUKKK")
			weather.Status.Water += 1
			weather.Status.Wind += 1
			fmt.Println("WATER => ", weather.Status.Water)
			fmt.Println("Wind => ", weather.Status.Wind)
			fmt.Println(strings.Repeat("-", 40))

			tpl, err := template.ParseFiles("template.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tpl.Execute(w, weather)
		})
		// return

		s.StartAsync()
		s.StartBlocking()

	}
}
