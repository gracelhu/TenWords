package main

//Problem I was running into before that I don't want to run into again:
//If you made changes to this rest api, you should go build first, not just ./introtoSWE,
//because the executable has an old version of the rest api saved. That's why postman
//will show inaccurate response sometimes

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
	"github.com/gen2brain/beeep"
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
var MapNametoPass = make(map[string]string)

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
	Tenwords   []Word `json:"tenwords"`
	StartIndex int    `json:"startindex"`
	Date       string `json:"date"` //in this format: 01-02-2006
}

// struct made for use in mongodb
type MongoField struct {
	ProgressIndex string            `json:"ProgressIndex"`
	Map           map[string]string `json:"map"`
}
type Auth struct {
	Username string            `json:"Username"`
	Password string            `json:"Password"`
	Date     string            `json:"date"` //in this format: 01-02-2006
	Map      map[string]string `json:"map"`
}
type AuthValidation struct {
	State string `json:"State"`
}
type QuizProgress struct {
	Username      string `json:"Username"`
	Quiz          string `json:"quiz"`
	QuestionCount string `json:"questioncount"`
}

/*
This function will be called by the route handler functions to fetch a word's information, like its:
definition, example sentence, audio file link, etc. This information is being fetched using a free
api called "Free Dictionary API". The struct used to demarshall the api's json response is in the file
"dictionaryapi.go". Additionally, this function will only be returning ENGLISH information. The information
will be translated to different languages in the route handler functions using the golang google translate api
*/

func setWordInfo(word *Word) {

	apiURL := dictionaryapiURL + word.Word

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

	word.English_definition = val[0].Meanings[0].Definitions[0].Definition
	for i := 0; i < len(val[0].Meanings); i++ {
		for x := 0; x < len(val[0].Meanings[i].Definitions); x++ {
			if val[0].Meanings[i].Definitions[x].Example != "" {
				word.Examplesentence_english = val[0].Meanings[i].Definitions[x].Example
			}
		}
	}

	for i := 0; i < len(val[0].Phonetics); i++ {
		if val[0].Phonetics[i].Audio != "" {
			word.Audiofilelink = val[0].Phonetics[i].Audio
		}
	}
}

func setTenWordsInfo(tenWords *TenWordPackage) {

	//array of all the api calls I want to make
	var apiURLs []string
	for i, _ := range tenWords.Tenwords {
		apiURLs = append(apiURLs, dictionaryapiURL+tenWords.Tenwords[i].Word)
	}

	// create a channel to store the responses
	respCh := make(chan string)

	// make GET requests concurrently
	for _, url := range apiURLs {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}
			defer resp.Body.Close()

			// read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body for %s: %v\n", url, err)
				return
			}

			// send the response body through the channel
			respCh <- string(body)
		}(url)
	}

	//Loop through the respChn
	for i := 0; i <= 9; i++ {
		resp := <-respCh

		var val []Words
		if err := json.Unmarshal([]byte(resp), &val); err != nil {
			panic(err)
		}

		tenWords.Tenwords[i].English_definition = val[0].Meanings[0].Definitions[0].Definition

		for i := 0; i < len(val[0].Meanings); i++ {
			for x := 0; x < len(val[0].Meanings[i].Definitions); x++ {
				if val[0].Meanings[i].Definitions[x].Example != "" {
					tenWords.Tenwords[i].Examplesentence_english = val[0].Meanings[i].Definitions[x].Example
				}
			}
		}

		for i := 0; i < len(val[0].Phonetics); i++ {
			if val[0].Phonetics[i].Audio != "" {
				tenWords.Tenwords[i].Audiofilelink = val[0].Phonetics[i].Audio
			}
		}
	}
}

func getTenWordsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ProgressIndexP = params["id"]
	updateWordProgress(params["id"])
	index, _ := strconv.Atoi(params["id"])

	/*
		Check to see if the tenwordpackage already exists in the allPackage object by looping through allPackage
		to see if a ten word package with params["id"] as its StartIndex exists or not. If yes, just set that object
		to vartenWords. If not, set all the parameters. This is a form of caching
	*/
	var noNeedToSet bool = false
	var tenWords TenWordPackage

	for i, _ := range allPackages {
		if allPackages[i].StartIndex == index {
			tenWords = allPackages[i]
			noNeedToSet = true
		}
	}

	if !noNeedToSet {
		//slices are zero indexed- not 1 indexed (starting from 0 not 1)
		tenWords = TenWordPackage{allWords[index-1 : index+9], index, currentTime.Format("01-02-2006")}

		setTenWordsInfo(&tenWords) //Passing in a word object by reference

		//I think that if you put a name for the second parameter (ex: tenworditem), it makes a COPY of the object, which is why I don't want to use it
		for i, _ := range tenWords.Tenwords {
			result, err = t.Translate(tenWords.Tenwords[i].Word, "auto", params["languagecode"])
			if err != nil {
				panic(err)
			}
			tenWords.Tenwords[i].Foreignword = result.Text

			result, _ = t.Translate(tenWords.Tenwords[i].English_definition, "auto", params["languagecode"])
			tenWords.Tenwords[i].Foreign_definition = result.Text

			result, _ = t.Translate(tenWords.Tenwords[i].Examplesentence_english, "auto", params["languagecode"])
			tenWords.Tenwords[i].Examplesentence_foreign = result.Text
		}
	}
	allPackages = append(allPackages, tenWords)
	json.NewEncoder(w).Encode(tenWords)
}

func getWord(w http.ResponseWriter, r *http.Request, words []Word,
	trans *translator.Translator) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through the words and find the one with the correct id
	for _, item := range words {
		if item.ID == params["id"] {
			transresult, _ := trans.Translate(item.Word, "auto", params["languagecode"])
			item.Foreignword = transresult.Text

			setWordInfo(&item)

			transresult, _ = trans.Translate(item.English_definition, "auto", params["languagecode"])
			item.Foreign_definition = transresult.Text
			transresult, _ = trans.Translate(item.Examplesentence_english, "auto", params["languagecode"])
			item.Examplesentence_foreign = transresult.Text

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

	//fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []MongoField
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	//for _, result := range results {
	//fmt.Printf("%+v\n", result)
	//}
	//end retrieve
}

// trying to make username and pass routehandler
func getnameandpass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	item := Auth{Username: params["username"], Password: params["password"], Date: dateP, Map: MapNametoPass}
	//item := Auth{Username: "Aeyesha", Password: "password123", Date: dateP, Map: MapNametoPass}
	var returnState = storeAuth(item)
	ret := AuthValidation{State: returnState}
	json.NewEncoder(w).Encode(ret)
}

// making database for username and pass
func storeAuth(auth Auth) string {
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
	col := client.Database("First_Database").Collection("AuthenticationDB")
	fmt.Println("Collection Type: ", reflect.TypeOf(col))
	MapNametoPass[auth.Username] = auth.Password
	//fmt.Println("printmap:", MapDatetoindex, "in:", ProgressIndexP)
	oneDoc := Auth{

		Username: auth.Username,
		Password: auth.Password,
		Date:     auth.Date,
		Map:      MapNametoPass,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc))
	//cutting and pasting new code
	//retrieve mongodb data
	filter := bson.M{"Username": auth.Username, "Password": auth.Password, "Date": auth.Date}
	update := bson.M{"$set": bson.M{
		"Username": auth.Username,
		"Password": auth.Password,
		"Date":     auth.Date,
	}}

	_, errp := col.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	//fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Auth
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result.Username)
		fmt.Printf("%+v\n", result.Password)
		fmt.Printf("%+v\n", result.Date)
	}
	var state string = "register"
	//var username bool = false
	for _, result := range results {
		//encode, _ := json.Marshal(result)
		if result.Username == auth.Username && result.Password == auth.Password {
			auth.Date = result.Date
			state = "returning"
			fmt.Println("Returning user!")
		} else if result.Username == auth.Username && result.Password != auth.Password {
			state = "invalid"
			fmt.Println("Invalid password error!")
		}

	}

	//end retrieve
	//end experiment
	if state == "register" {
		result, insertErr := col.InsertOne(ctx, oneDoc)
		fmt.Printf(results[len(results)-1].Date)
		auth.Date = results[len(results)-1].Date
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
	}

	//fmt.Println(auth.Date);

	return state + "|" + auth.Date

	//end mongo

}
func getquizprogress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
	/*
		err := beeep.Alert("TenWords", "Great job on finishing the quiz! You are ready to learn ten new words!", "")
		if err != nil {
			panic(err)
		}
	*/
	fmt.Printf("Test running quiz progress")
	item := QuizProgress{Username: params["username"], Quiz: params["quiznumber"], QuestionCount: params["questioncount"]}
	storeQuiz(item)
	//storeAuth(item)
	json.NewEncoder(w).Encode(item)
}

func storeQuiz(quiz QuizProgress) {
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
	col := client.Database("First_Database").Collection("QuizDB")
	fmt.Println("Collection Type: ", reflect.TypeOf(col))

	oneDoc := QuizProgress{

		Username: quiz.Username,
		Quiz:     quiz.Quiz,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc))
	//cutting and pasting new code
	//retrieve mongodb data
	filter := bson.M{"Username": quiz.Username, "Quiz": quiz.Quiz}
	update := bson.M{"$set": bson.M{
		"Username": quiz.Username,
		"quiz":     quiz.Quiz,
	}}

	_, errp := col.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []QuizProgress
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

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

	//Route Handler for fetching a single word in each of the foreign languages by their index
	//Need to pass in allWords slice and t translator variable into route handler functions, orelse unit
	//tests treats those variables as not existing
	r.HandleFunc("/api/words/{languagecode}/single/{id}", func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, allWords, t)
	})

	//Route Handler for fetching 10 word packages in each of the foreign languages by their date
	r.HandleFunc("/api/words/{languagecode}/package/date/{date}", getTenWordsByDate).Methods("GET")

	//Route Handler for authentication
	//r.HandleFunc("/api/words/package/auth", getnameandpass).Methods("GET")
	r.HandleFunc("/auth/{username}/{password}", getnameandpass).Methods("GET")

	r.HandleFunc("/quiz/{username}/{quiznumber}/{questioncount}", getquizprogress).Methods("GET")

	//r.HandleFunc("/startdate/{username}/{quiznumber}", getstartdate).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

	resp, err := http.Get("/api/words/es/package/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
