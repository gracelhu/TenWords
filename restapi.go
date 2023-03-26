package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	translator "github.com/Conight/go-googletrans"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var allWords []Word                                                       //stores all words
var allPackages []TenWordPackage                                          //stores all the ten word packages created
var currentTime = time.Now()                                              //to get current date and time
var t = translator.New()                                                  //Init translator
var result, err = t.Translate("", "auto", "en")                           //Init result of translation
var dictionaryapiURL = "https://api.dictionaryapi.dev/api/v2/entries/en/" //the word we're interested in fetching information for will be appended to this url
// variables for params["id"], date, and map use in Mongodb
var dateP = currentTime.Format("01-02-2006")
var ProgressIndexP string
var MapDatetoindex = make(map[string]string)

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

type VocabWord struct {
	Tenwords []Word `json:"tenvocabwords"`
}

type TenWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}

// only includes an array of the English words for the ten word package. This is for fast display of the
// "preview" card (react mui card component with packet #, date, and english words on it)
type TenWordVocabPackage struct {
	TenwordsVocab []Word `json:"tenwordsvocab"`
}

// struct made for use in mongodb
type MongoField struct {
	ProgressIndex string            `json:"ProgressIndex"`
	Map           map[string]string `json:"map"`
}

/*
This function will be called by the route handler functions to fetch a word's information, like its:
definition, example sentence, audio file link, etc. This information is being fetched using a free
api called "Free Dictionary API". The struct used to demarshall the api's json response is in the file
"dictionaryapi.go". Additionally, this function will only be returning ENGLISH information. The information
will be translated to different languages in the route handler functions using the golang google translate api
*/

func getWordInfo(word string, infoType string) string {

	apiURL := dictionaryapiURL + word

	resp, err := http.Get(apiURL)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
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
func getTenWordsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ProgressIndexP = params["id"]
	updateWordProgress(params["id"])
	index, _ := strconv.Atoi(params["id"])
	var tenWords = TenWordPackage{allWords[index : index+10], currentTime.Format("01-02-2006")}
	for i, tenworditem := range tenWords.Tenwords {

		//Set the English_definition, Examplesentence_english, and Audiofilelink fields by calling
		//the function getWordInfo
		tenWords.Tenwords[i].English_definition = getWordInfo(tenWords.Tenwords[i].Word, "definition")
		tenWords.Tenwords[i].Audiofilelink = getWordInfo(tenWords.Tenwords[i].Word, "audiofilelink")
		tenWords.Tenwords[i].Examplesentence_english = getWordInfo(tenWords.Tenwords[i].Word, "examplesentence")
		result, err = t.Translate(tenworditem.Word, "auto", params["languagecode"])
		if err != nil {
			panic(err)
		}
		tenWords.Tenwords[i].Foreignword = result.Text

		result, _ = t.Translate(tenWords.Tenwords[i].English_definition, "auto", params["languagecode"])
		tenWords.Tenwords[i].Foreign_definition = result.Text

		result, _ = t.Translate(tenWords.Tenwords[i].Examplesentence_english, "auto", params["languagecode"])
		tenWords.Tenwords[i].Examplesentence_foreign = result.Text
	}
	allPackages = append(allPackages, tenWords)
	json.NewEncoder(w).Encode(tenWords)
}

func getTenWordsVocabByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	index, _ := strconv.Atoi(params["id"])

	var tenWordsVocab = TenWordVocabPackage{allWords[index : index+10]}
	json.NewEncoder(w).Encode(tenWordsVocab)
	json.NewEncoder(w).Encode(&TenWordVocabPackage{})
}

func getWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	for _, item := range allWords {
		if item.ID == params["id"] {
			result, err = t.Translate(item.Word, "auto", params["languagecode"])
			if err != nil {
				panic(err)
			}

			item.Foreignword = result.Text
			item.English_definition = getWordInfo(item.Word, "definition")
			item.Audiofilelink = getWordInfo(item.Word, "audiofilelink")
			item.Examplesentence_english = getWordInfo(item.Word, "examplesentence")

			result, _ = t.Translate(item.English_definition, "auto", params["languagecode"])
			item.Foreign_definition = result.Text
			result, _ = t.Translate(item.Examplesentence_english, "auto", params["languagecode"])
			item.Examplesentence_foreign = result.Text

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Word{})
}

// this doesn't really work right now but I'll fix it later
func getTenWordsByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the allPackages slice and find the one with the correct date
	for _, item := range allPackages {
		if item.Date == params["date"] {
			for i, tenworditem := range item.Tenwords {

				result, err = t.Translate(tenworditem.Word, "auto", params["languagecode"])
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

// Function that updates index after each api call
func updateWordProgress(progressIndex string) {
	//mongo stuff
	//Pushing data to mongodb
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, call := context.WithTimeout(context.Background(), 15*time.Second)
	defer call()
	col := client.Database("First_Database").Collection("First COllection3")
	fmt.Println("Collection Type: ", reflect.TypeOf(col))

	//MapDatetoindex[dateP] = ProgressIndexP
	/*So right now, when we're testing the api, we won't be making api calls on seperate days.
	Therefore, since there can't be repeat days, the map will only store the first call made that day.
	To work around this for testing purposes, we'll just store the ProgressIndexP as the key too */

	MapDatetoindex[dateP] = ProgressIndexP
	fmt.Println("printmap:", MapDatetoindex, "in:", ProgressIndexP)
	oneDoc := MongoField{

		ProgressIndex: ProgressIndexP,
		Map:           MapDatetoindex,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc))

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}

	//end mongo
	//retrieve mongodb data
	filter := bson.M{"ProgressIndex": ProgressIndexP}
	update := bson.M{"$set": bson.M{
		"progressindex": progressIndex,
	}}

	_, errp := col.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []MongoField
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
	//end retrieve
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

	//Route Handler for fetching 10 word packages in each of the foreign languages by their starting index
	r.HandleFunc("/api/words/{languagecode}/package/{id}", getTenWordsByID).Methods("GET")

	//Route Handler for fetching 10 word packages (just the english vocab) in each of the foreign languages by their starting index
	//Why is this route handler not working..???
	r.HandleFunc("/api/words/package/vocab/{id}", getTenWordsVocabByID).Methods("GET")

	//Route Handler for fetching a single word in each of the foreign languages by their index
	r.HandleFunc("/api/words/{languagecode}/single/{id}", getWord).Methods("GET")

	//Route Handler for fetching 10 word packages in each of the foreign languages by their date
	r.HandleFunc("/api/words/{languagecode}/package/date/{date}", getTenWordsByDate).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

/*w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r)

//Loop through the words and find the one with the correct id
//the correct id for ten word packets: scalar * 10 + 1 --> 1, 11, 21, 31, 41, etc.

for index, item := range allWords {
	if item.ID == params["id"] {
		ProgressIndexP = params["id"]
		updateWordProgress(params["id"])
		var tenWords = TenWordPackage{allWords[index : index+10], currentTime.Format("01-02-2006")}
		for i, tenworditem := range tenWords.Tenwords {

			//Set the English_definition, Examplesentence_english, and Audiofilelink fields by calling
			//the function getWordInfo
			tenWords.Tenwords[i].English_definition = getWordInfo(tenWords.Tenwords[i].Word, "definition")
			tenWords.Tenwords[i].Audiofilelink = getWordInfo(tenWords.Tenwords[i].Word, "audiofilelink")
			tenWords.Tenwords[i].Examplesentence_english = getWordInfo(tenWords.Tenwords[i].Word, "examplesentence")
			result, err = t.Translate(tenworditem.Word, "auto", params["languagecode"])
			if err != nil {
				panic(err)
			}
			tenWords.Tenwords[i].Foreignword = result.Text

			result, _ = t.Translate(tenWords.Tenwords[i].English_definition, "auto", params["languagecode"])
			tenWords.Tenwords[i].Foreign_definition = result.Text

			result, _ = t.Translate(tenWords.Tenwords[i].Examplesentence_english, "auto", params["languagecode"])
			tenWords.Tenwords[i].Examplesentence_foreign = result.Text
		}
		allPackages = append(allPackages, tenWords)
		json.NewEncoder(w).Encode(tenWords)
		return
	}
} */
