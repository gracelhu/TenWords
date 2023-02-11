//A functional CRUD RESTFUL API for TenWords application

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var allWords []Word              //stores all words
var allPackages []TenWordPackage //stores all the ten word packages created
var currentTime = time.Now()     //to get current date and time

type Word struct {
	ID              string `json:"id"`
	Word            string `json:"word"`
	Spanish         string `json:"spanish"`
	Examplesentence string `json:"examplesentence"`
}

type TenWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}

// for small scale testing
type TwoWordPackage struct {
	Tenwords []Word `json:"twowords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}

func getAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allWords)
}

func getTenWordsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.
	//the correct id for two word packets: scalar * 10 + 1 --> 1, 3, 5, 7, etc.

	for index, item := range allWords {
		if item.ID == params["id"] {
			var tenWords = TenWordPackage{allWords[index : index+10], currentTime.Format("01-02-2006")}
			allPackages = append(allPackages, tenWords)
			json.NewEncoder(w).Encode(tenWords)
			return
		}
	}
}

func getTenWordsByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the allPackages slice and find the one with the correct date
	for _, item := range allPackages {
		if item.Date == params["date"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// for small scale testing
func getTwoWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.
	//the correct id for two word packets: scalar * 10 + 1 --> 1, 3, 5, 7, etc.

	for index, item := range allWords {
		if item.ID == params["id"] {
			var twoWords = TwoWordPackage{allWords[index : index+2], currentTime.Format("01-02-2006")}
			json.NewEncoder(w).Encode(twoWords)
			return
		}
	}
}

func getWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	for _, item := range allWords {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Word{})
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//@todo - implement DB
	//Take all words from wordlist.txt and put them in allWords slice
	readFile, err := os.Open("wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	index := 1
	for fileScanner.Scan() {
		allWords = append(allWords, Word{ID: strconv.Itoa(index), Word: fileScanner.Text()})
		index++
	}

	//Route Handlers
	r.HandleFunc("/api/words", getAllWords).Methods("GET")
	r.HandleFunc("/api/words/package/id/{id}", getTenWordsByID).Methods("GET")
	r.HandleFunc("/api/words/single/{id}", getWord).Methods("GET")
	r.HandleFunc("/api/words/package/date/{date}", getTenWordsByDate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
