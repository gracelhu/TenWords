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

	translator "github.com/Conight/go-googletrans"
	"github.com/gorilla/mux"
)

var allWords []Word                             //stores all words
var allPackages []TenWordPackage                //stores all the ten word packages created
var currentTime = time.Now()                    //to get current date and time
var t = translator.New()                        //Init translator
var result, err = t.Translate("", "auto", "en") //Init result of translation

type Word struct {
	ID              string `json:"id"`
	Word            string `json:"english"`  //en
	Spanish         string `json:"spanish"`  //es
	French          string `json:"french"`   //fr
	Russian         string `json:"russian"`  //ru
	Italian         string `json:"italian"`  //it
	Japanese        string `json:"japanese"` //ja
	Chinese         string `json:"chinese"`  //zh-cn
	Examplesentence string `json:"examplesentence"`
}

type TenWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}

func getTenWordsByID(w http.ResponseWriter, r *http.Request, languageCode string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.

	for index, item := range allWords {
		if item.ID == params["id"] {
			var tenWords = TenWordPackage{allWords[index : index+10], currentTime.Format("01-02-2006")}
			for i, tenworditem := range tenWords.Tenwords {

				result, err = t.Translate(tenworditem.Word, "auto", languageCode)
				if err != nil {
					panic(err)
				}

				if languageCode == "es" {
					tenWords.Tenwords[i].Spanish = result.Text
				}
				if languageCode == "fr" {
					tenWords.Tenwords[i].French = result.Text
				}
				if languageCode == "ru" {
					tenWords.Tenwords[i].Russian = result.Text
				}
				if languageCode == "it" {
					tenWords.Tenwords[i].Italian = result.Text
				}
				if languageCode == "ja" {
					tenWords.Tenwords[i].Japanese = result.Text
				}
				if languageCode == "zh-cn" {
					tenWords.Tenwords[i].Chinese = result.Text
				}

			}
			allPackages = append(allPackages, tenWords)
			json.NewEncoder(w).Encode(tenWords)
			return
		}
	}
}

func getWord(w http.ResponseWriter, r *http.Request, languageCode string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	for _, item := range allWords {
		if item.ID == params["id"] {
			result, err := t.Translate(item.Word, "auto", languageCode)
			if err != nil {
				panic(err)
			}

			//Setting the language fields
			if languageCode == "es" {
				item.Spanish = result.Text
			}
			if languageCode == "fr" {
				item.French = result.Text
			}
			if languageCode == "ru" {
				item.Russian = result.Text
			}
			if languageCode == "it" {
				item.Italian = result.Text
			}
			if languageCode == "ja" {
				item.Japanese = result.Text
			}
			if languageCode == "zh-cn" {
				item.Chinese = result.Text
			}

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Word{})
}

// this doesn't really work right now but I'll fix it later
func getTenWordsByDate(w http.ResponseWriter, r *http.Request, languageCode string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the allPackages slice and find the one with the correct date
	for _, item := range allPackages {
		if item.Date == params["date"] {
			for i, tenworditem := range item.Tenwords {

				result, err = t.Translate(tenworditem.Word, "auto", languageCode)
				if err != nil {
					panic(err)
				}

				if languageCode == "es" {
					item.Tenwords[i].Spanish = result.Text
				}
				if languageCode == "fr" {
					item.Tenwords[i].French = result.Text
				}
				if languageCode == "ru" {
					item.Tenwords[i].Russian = result.Text
				}
				if languageCode == "it" {
					item.Tenwords[i].Italian = result.Text
				}
				if languageCode == "ja" {
					item.Tenwords[i].Japanese = result.Text
				}
				if languageCode == "zh-cn" {
					item.Tenwords[i].Chinese = result.Text
				}
			}
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //Init Router

	//Take all words from wordlist.txt and put them in allWords slice
	readFile, err := os.Open("wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	index := 1
	for fileScanner.Scan() {
		newWord := Word{ID: strconv.Itoa(index), Word: fileScanner.Text()}
		allWords = append(allWords, newWord)
		index++
	}

	/* ===== Route Handlers ===== */

	//Route Handlers for fetching 10 word packages in each of the foreign languages
	r.HandleFunc("/api/words/spanish/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "es")
	}).Methods("GET")
	r.HandleFunc("/api/words/french/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "fr")
	}).Methods("GET")
	r.HandleFunc("/api/words/russian/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "ru")
	}).Methods("GET")
	r.HandleFunc("/api/words/italian/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "it")
	}).Methods("GET")
	r.HandleFunc("/api/words/japanese/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "ja")
	}).Methods("GET")
	r.HandleFunc("/api/words/chinese/package/{id}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "zh-cn")
	}).Methods("GET")

	//Route Handlers for fetching a single word in each of the foreign languages (for testing)
	r.HandleFunc("/api/words/spanish/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "es")
	}).Methods("GET")
	r.HandleFunc("/api/words/french/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "fr")
	}).Methods("GET")
	r.HandleFunc("/api/words/russian/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "ru")
	}).Methods("GET")
	r.HandleFunc("/api/words/italian/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "it")
	}).Methods("GET")
	r.HandleFunc("/api/words/japanese/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "ja")
	}).Methods("GET")
	r.HandleFunc("/api/words/chinese/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, "zh-cn")
	}).Methods("GET")

	//Route Handlers for fetching 10 word packages in each of the foreign languages by their date
	r.HandleFunc("/api/words/spanish/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "es")
	}).Methods("GET")
	r.HandleFunc("/api/words/french/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "fr")
	}).Methods("GET")
	r.HandleFunc("/api/words/russian/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "ru")
	}).Methods("GET")
	r.HandleFunc("/api/words/italian/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "it")
	}).Methods("GET")
	r.HandleFunc("/api/words/japanese/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "ja")
	}).Methods("GET")
	r.HandleFunc("/api/words/chinese/package/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		getTenWordsByID(w, r, "zh-cn")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//See this: https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler
