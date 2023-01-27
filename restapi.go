//A functional CRUD RESTFUL API for TenWords application
//Test and make API calls with Postman 

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var allWords []Word //stores all words

type Word struct {
	ID              string `json:"id"`
	Word            string `json:"word"`
	Definition      string `json:"definition"`
	Partofspeech    string `json:"partofspeech"`
	Examplesentence string `json:"examplesentence"`
}

// fetching two words for now to make it easier to test, will fetch 10 words later
type TwoWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"`
}

// Get All Books
func getAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allWords)
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
			var twoWords = TwoWordPackage{allWords[index : index+2], "1/27/23"}
			json.NewEncoder(w).Encode(twoWords)
			return
		}
	}
}

func getTenWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.
	//the correct id for two word packets: scalar * 10 + 1 --> 1, 3, 5, 7, etc.

	for index, item := range allWords {
		if item.ID == params["id"] {
			var twoWords = TwoWordPackage{allWords[index : index+10], "1/27/23"}
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

	//Mock Data - @todo - implement DB
	allWords = append(allWords, Word{ID: "1", Word: "hello"})
	allWords = append(allWords, Word{ID: "2", Word: "apple"})
	allWords = append(allWords, Word{ID: "3", Word: "bye"})
	allWords = append(allWords, Word{ID: "4", Word: "how"})
	allWords = append(allWords, Word{ID: "5", Word: "blue"})
	allWords = append(allWords, Word{ID: "6", Word: "red"})
	allWords = append(allWords, Word{ID: "7", Word: "eyes"})
	allWords = append(allWords, Word{ID: "8", Word: "nose"})
	allWords = append(allWords, Word{ID: "9", Word: "running"})
	allWords = append(allWords, Word{ID: "10", Word: "awesome"})

	//Route Handlers
	r.HandleFunc("/api/words", getAllWords).Methods("GET")
	r.HandleFunc("/api/words/package/{id}", getTenWords).Methods("GET")
	r.HandleFunc("/api/words/single/{id}", getWord).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
