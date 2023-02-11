/*
	- A functional CRUD RESTFUL API for TenWords application
	- This API fetches ten word packages from a text file called "wordlist.txt" and uses
	  golang googletrans library to translate the words into 6 different foreign languages:
	  spanish, french, russian, italian, japanese, and chinese.
	- Every time a ten word package is fetched, the date that package was fetched is recorded and
	  stored in our database
	- Additionally, it stores the language learning progress of users in a database by recording
	  what index in wordlist.txt the user left off on.
	- This API also autogenerates example sentences for each vocabulary word by calling a repository
	  of millions of english sentences stored in a database
*/

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

var allWords []Word              //stores all words
var allPackages []TenWordPackage //stores all the ten word packages created
var currentTime = time.Now()     //to get current date and time
var t = translator.New()         //Init translator

// Our application lets users learn these languages: spanish, french, russian, italian, japanese, chinese (simplified)
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

func getAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allWords)
}

func getTenWordsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.

	for index, item := range allWords {
		if item.ID == params["id"] {
			var tenWords = TenWordPackage{allWords[index : index+10], currentTime.Format("01-02-2006")}
			for i, _ := range tenWords.Tenwords {
				languageCodes := [6]string{"es", "fr", "ru", "it", "ja", "zh-cn"}
				for _, languageCode := range languageCodes {

					result, err := t.Translate(item.Word, "auto", languageCode)
					if err != nil {
						panic(err)
					}

					//Setting the language fields
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
			}
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

func getWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	for _, item := range allWords {
		if item.ID == params["id"] {
			languageCodes := [6]string{"es", "fr", "ru", "it", "ja", "zh-cn"}

			for _, languageCode := range languageCodes {
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
			}

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Word{})
}

/*func setForeignLanguageFields(item Word*) {

} */

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
		newWord := Word{ID: strconv.Itoa(index), Word: fileScanner.Text()}
		allWords = append(allWords, newWord)
		index++
	}

	//Route Handlers
	r.HandleFunc("/api/words", getAllWords).Methods("GET")
	r.HandleFunc("/api/words/package/id/{id}", getTenWordsByID).Methods("GET")
	r.HandleFunc("/api/words/single/{id}", getWord).Methods("GET")
	r.HandleFunc("/api/words/package/date/{date}", getTenWordsByDate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
