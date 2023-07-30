package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID           string    `json:"id"`
	ISBN         string    `json:"isbn"`
	Title        string    `json:"title"`
	Release_Year string    `json:"release_year"`
	Director     *Director `json:"director"`
}

type Director struct {
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "41ds1cs54", Title: "Dabbang", Release_Year: "2006", Director: &Director{First_Name: "Anurag", Last_Name: "Kashyap"}})

	movies = append(movies, Movie{ID: "2", ISBN: "jsfbd5f4d", Title: "Rockstar", Release_Year: "2010", Director: &Director{First_Name: "Imtiaz", Last_Name: "Ali"}})

	router.HandleFunc("/getallmovies", GetMovies).Methods("GET")
	router.HandleFunc("/getmovie/{id}", GetMovie).Methods("GET")
	router.HandleFunc("/deletemovie/{id}", DeleteMovie).Methods("GET")
	router.HandleFunc("/createmovie", CreateMovie).Methods("POST")
	router.HandleFunc("/updatemovie/{id}", UpdateMovie).Methods("POST")

	fmt.Println("Server is going to listen at port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func GetMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func GetMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range movies {
		if params["id"] == item.ID {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func CreateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie

	json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movie)
}

func UpdateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie

	json.NewDecoder(req.Body).Decode(&movie)
	params := mux.Vars(req)

	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)

			movie.ID = params["id"]
			movies = append(movies, movie)
			break
		}
	}
	json.NewEncoder(res).Encode(movie)

}
