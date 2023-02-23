package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	translator "github.com/Conight/go-googletrans"
	"github.com/gorilla/mux"
)

var allWords []Word                                                       //stores all words
var allPackages []TenWordPackage                                          //stores all the ten word packages created
var currentTime = time.Now()                                              //to get current date and time
var t = translator.New()                                                  //Init translator
var result, err = t.Translate("", "auto", "en")                           //Init result of translation
var dictionaryapiURL = "https://api.dictionaryapi.dev/api/v2/entries/en/" //the word we're interested in fetching information for will be appended to this url

type Word struct {
	ID                      string `json:"id"`
	Word                    string `json:"english"`     //en
	Foreignword             string `json:"foreignword"` //could be es, fr, ru, it, ja, or zh-cn
	Examplesentence_english string `json:"examplesentence_english"`
	Examplesentence_foreign string `json:"examplesentence_foreign"`
	English_definition      string `json:"english_definition"`
	Foreign_definition      string `json:"foreign_definition"`
	Audiofilelink           string `json:"audiofilelink"`
}

type TenWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}

// This function will be called by the route handler functions to fetch a word's information, like its:
// definition, example sentence, audio file link, etc. This information is being fetched using a free
// api called "Free Dictionary API". The struct used to demarshall the api's json response is in the file
// "dictionaryapi.go". Additionally, this function will only be returning ENGLISH information. The information
// will be translated to different languages in the route handler functions using the golang google translate api

func getWordInfo(word string, infoType string) string {

	apiURL := dictionaryapiURL + word

	resp, err := http.Get(apiURL)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var val []Words
	if err := json.Unmarshal([]byte(body), &val); err != nil {
		panic(err)
	}

	if infoType == "definition" {
		return val[0].Meanings[0].Definitions[0].Definition
	}
	if infoType == "examplesentence" {
		for i := 0; i < len(val[0].Meanings); i++ {
			for x := 0; x < len(val[0].Meanings[i].Definitions); x++ {
				if val[0].Meanings[i].Definitions[x].Example != "" {
					return val[0].Meanings[i].Definitions[x].Example
				}
			}
		}
	}
	if infoType == "audiofilelink" {
		for i := 0; i < len(val[0].Phonetics); i++ {
			if val[0].Phonetics[i].Audio != "" {
				return val[0].Phonetics[i].Audio
			}
		}
	}
	return ""
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

				//Set the English_definition, Examplesentence_english, and Audiofilelink fields by calling
				//the function getWordInfo
				tenWords.Tenwords[i].English_definition = getWordInfo(tenWords.Tenwords[i].Word, "definition")
				tenWords.Tenwords[i].Audiofilelink = getWordInfo(tenWords.Tenwords[i].Word, "audiofilelink")
				tenWords.Tenwords[i].Examplesentence_english = getWordInfo(tenWords.Tenwords[i].Word, "examplesentence")
				result, err = t.Translate(tenworditem.Word, "auto", languageCode)
				if err != nil {
					panic(err)
				}
				tenWords.Tenwords[i].Foreignword = result.Text

				result, _ = t.Translate(tenWords.Tenwords[i].English_definition, "auto", languageCode)
				tenWords.Tenwords[i].Foreign_definition = result.Text

				result, _ = t.Translate(tenWords.Tenwords[i].Examplesentence_english, "auto", languageCode)
				tenWords.Tenwords[i].Examplesentence_foreign = result.Text
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
			item.Foreignword = result.Text

			result, _ = t.Translate(item.English_definition, "auto", languageCode)
			item.Foreign_definition = result.Text

			result, _ = t.Translate(item.Examplesentence_english, "auto", languageCode)
			item.Examplesentence_foreign = result.Text

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
				item.Tenwords[i].Foreignword = result.Text
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
