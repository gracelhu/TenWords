package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"

	"testing"

	translator "github.com/Conight/go-googletrans"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type expectedWordResponse struct {
	ID                     string `json:"id"`
	English                string `json:"english"`
	ForeignWord            string `json:"foreignword"`
	ExampleSentenceEnglish string `json:"examplesentence_english"`
	ExampleSentenceForeign string `json:"examplesentence_foreign"`
	EnglishDefinition      string `json:"english_definition"`
	ForeignDefinition      string `json:"foreign_definition"`
	AudioFileLink          string `json:"audiofilelink"`
}

type expectedTenWordResponse struct {
	Tenwords []expectedWordResponse `json:"tenwords"`
	Date     string                 `json:"date"`
}

var mockWords = []Word{
	{ID: "1", Word: "abandon"},
	{ID: "2", Word: "sudden"},
	{ID: "3", Word: "lawyer"},
	{ID: "4", Word: "particularly"},
	{ID: "5", Word: "gender"},
	{ID: "6", Word: "literary"},
	{ID: "7", Word: "cotton"},
	{ID: "8", Word: "station"},
	{ID: "9", Word: "everyone"},
	{ID: "10", Word: "life"},
}

func TestUpdateWordProgress(t *testing.T) {
	// Connect to the test database
	def := "testindex"
	//updated:="testindex2"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			t.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Create a new collection for testing
	testColl := client.Database("testdb").Collection("testcoll")

	// Insert a test document into the collection
	doc := MongoField{
		ProgressIndex: def,
		Map:           map[string]string{"testkey": "testvalue"},
	}
	_, err = testColl.InsertOne(context.Background(), doc)
	if err != nil {
		t.Fatal("Failed to insert test document:", err)
	}

	// Call the function with the test index
	//updateWordProgress("testindex2")
	//def="testindex2"
	// Check that the index was updated in the database
	filter := bson.M{"ProgressIndex": def}

	update := bson.M{"$set": bson.M{
		"progressindex": def,
	}}

	_, errp := testColl.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	//fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := testColl.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []MongoField
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
		if result.ProgressIndex != def {
			fmt.Println("res ", result.ProgressIndex)
			t.Fatal("Update failed: progress index not updated in database")
		}
	}

}
func TestWord(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			t.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Create a new collection for testing
	testColl := client.Database("testdb").Collection("testcoll2")

	doc := Word{

		ID:                      "6",
		Word:                    "literary",
		Foreignword:             "literary",
		Examplesentence_english: "a literary history",
		Examplesentence_foreign: "a literary history",
		English_definition:      "Relating to literature.",
		Foreign_definition:      "Relating to literature.",
		Audiofilelink:           "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3",
	}
	_, err = testColl.InsertOne(context.Background(), doc)
	if err != nil {
		t.Fatal("Failed to insert test document:", err)
	}
	cursor, err := testColl.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	filter := bson.M{"ID": "6"}

	update := bson.M{"$set": bson.M{

		"id": "6",
	}}

	_, errp := testColl.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}
	var results []Word
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
		if result.ID != "6" {
			fmt.Println("res ", result.ID)
			t.Fatal("id not equal")
		}
		if result.Word != "literary" {
			fmt.Println("res ", result.Word)
			t.Fatal("word not equal")
		}
		if result.Foreignword != "literary" {
			fmt.Println("res ", result.Foreignword)
			t.Fatal("forword not equal")
		} else {

			fmt.Println("test passed")
		}
		if result.Examplesentence_english != "a literary history" {
			fmt.Println("res ", result.Examplesentence_english)
			t.Fatal("engsen not equal")
		}
		if result.Examplesentence_foreign != "a literary history" {
			fmt.Println("res ", result.Examplesentence_foreign)
			t.Fatal("forsen not equal")
		}
		if result.English_definition != "Relating to literature." {
			fmt.Println("res ", result.English_definition)
			t.Fatal("engdef not equal")
		}
		if result.Foreign_definition != "Relating to literature." {
			fmt.Println("res ", result.Foreign_definition)
			t.Fatal("fordef not equal")
		}
		if result.Audiofilelink != "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3" {
			fmt.Println("res ", result.Audiofilelink)
			t.Fatal("audio not equal")
		}

	}
}

func TestUsername(t *testing.T) {
	// Connect to the test database
	def := "Aayesha"
	//updated:="testindex2"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			t.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Create a new collection for testing
	testColl := client.Database("testdb").Collection("testcoll3")

	// Insert a test document into the collection

	doc := Auth{

		Username: "Aayesha",
		Password: "132",
		Date:     dateP,
		Map:      MapNametoPass,
	}

	_, err = testColl.InsertOne(context.Background(), doc)
	if err != nil {
		t.Fatal("Failed to insert test document:", err)
	}

	// Call the function with the test index
	//updateWordProgress("testindex2")
	//def="testindex2"
	// Check that the index was updated in the database
	filter := bson.M{"Username": "Aayesha", "Password": "132", "Date": dateP}
	update := bson.M{"$set": bson.M{
		"Username": "Aayesha",
		"password": "132", "Date": dateP,
	}}

	_, errp := testColl.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	//fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := testColl.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Auth
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
		if result.Username != def {

			t.Fatal("Update failed: username not updated in database")
		}
	}

}

func TestPassword(t *testing.T) {
	// Connect to the test database
	def := "132"
	//updated:="testindex2"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			t.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Create a new collection for testing
	testColl := client.Database("testdb").Collection("testcoll3")

	// Insert a test document into the collection

	doc := Auth{

		Username: "Aayesha",
		Password: "132",
		Date:     dateP,
		Map:      MapNametoPass,
	}

	_, err = testColl.InsertOne(context.Background(), doc)
	if err != nil {
		t.Fatal("Failed to insert test document:", err)
	}

	// Call the function with the test index
	//updateWordProgress("testindex2")
	//def="testindex2"
	// Check that the index was updated in the database
	filter := bson.M{"Username": "Aayesha", "Password": "132", "Date": dateP}
	update := bson.M{"$set": bson.M{
		"Username": "Aayesha",
		"password": "132", "Date": dateP,
	}}

	_, errp := testColl.UpdateOne(context.Background(), filter, update)
	if errp != nil {
		log.Fatal(err)
	}

	//fmt.Println("Word progress updated to: ", progressIndex)
	cursor, err := testColl.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Auth
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
		if result.Password != def {

			t.Fatal("Update failed: username not updated in database")
		}
	}

}

func TestGetChineseWord(t *testing.T) {
	var trans = translator.New() //Init translator

	// Create a closure that wraps the getWord function
	handler := func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, mockWords, trans)
	}

	// Set up the request
	req, err := http.NewRequest("GET", "/api/words/zh/single/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the response recorder
	rr := httptest.NewRecorder()

	// Set up the router and register the closure as the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/single/{id}", handler)

	// Make the request
	router.ServeHTTP(rr, req)

	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected expectedWordResponse
	if err := json.Unmarshal([]byte(`{
		"id": "2",
		"english": "sudden",
		"foreignword": "突然的",
		"examplesentence_english": "The sudden drop in temperature left everyone cold and confused.",
		"examplesentence_foreign": "突如其来的降温，让所有人都感到寒冷和迷茫。",
		"english_definition": "An unexpected occurrence; a surprise.",
		"foreign_definition": "意外事件；惊喜。",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3"
	}`), &expected); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	var actual expectedWordResponse
	if err := json.Unmarshal([]byte(rr.Body.String()), &actual); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %+v want %+v", actual, expected)
	}
}

func TestGetSpanishWord(t *testing.T) {
	var trans = translator.New() //Init translator

	// Create a closure that wraps the getWord function
	handler := func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, mockWords, trans)
	}

	// Set up the request
	req, err := http.NewRequest("GET", "/api/words/es/single/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the response recorder
	rr := httptest.NewRecorder()

	// Set up the router and register the closure as the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/single/{id}", handler)

	// Make the request
	router.ServeHTTP(rr, req)

	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected expectedWordResponse
	if err := json.Unmarshal([]byte(`{
		"id": "1",
		"english": "abandon",
		"foreignword": "abandonar",
		"examplesentence_english": "Many baby girls have been abandoned on the streets of Beijing.",
		"examplesentence_foreign": "Muchas niñas han sido abandonadas en las calles de Beijing.",
		"english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.",
		"foreign_definition": "Renunciar o renunciar al control de, rendirse o entregarse, o ceder a las propias emociones.",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3"
	}`), &expected); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	var actual expectedWordResponse
	if err := json.Unmarshal([]byte(rr.Body.String()), &actual); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %+v want %+v", actual, expected)
	}
}

func TestTenWordsByID(t *testing.T) {
	var trans = translator.New() //Init translator

	// Create a closure that wraps the getWord function
	handler := func(w http.ResponseWriter, r *http.Request) {
		getWord(w, r, mockWords, trans)
	}

	// Set up the request
	req, err := http.NewRequest("GET", "/api/words/it/package/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the response recorder
	rr := httptest.NewRecorder()

	// Set up the router and register the closure as the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/package/{id}", handler)

	// Make the request
	router.ServeHTTP(rr, req)

	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected expectedTenWordResponse
	if err := json.Unmarshal([]byte(`{
		"tenwords": [
			{
				"id": "1",
				"english": "abandon",
				"foreignword": "abbandono",
				"examplesentence_english": "Many baby girls have been abandoned on the streets of Beijing.",
				"examplesentence_foreign": "Molte bambine sono state abbandonate per le strade di Pechino.",
				"english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.",
				"foreign_definition": "Rinunciare o rinunciare al controllo, arrendersi o arrendersi, o cedere alle proprie emozioni.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3"
			},
			{
				"id": "2",
				"english": "sudden",
				"foreignword": "improvviso",
				"examplesentence_english": "The sudden drop in temperature left everyone cold and confused.",
				"examplesentence_foreign": "L'improvviso calo della temperatura ha lasciato tutti infreddoliti e confusi.",
				"english_definition": "An unexpected occurrence; a surprise.",
				"foreign_definition": "Un evento inaspettato; una sorpresa.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3"
			},
			{
				"id": "3",
				"english": "lawyer",
				"foreignword": "avvocato",
				"examplesentence_english": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation",
				"examplesentence_foreign": "Il tempo e i consigli di un avvocato sono il suo mestiere. - aforisma spesso accreditato ad Abraham Lincoln, ma senza attestazione",
				"english_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.",
				"foreign_definition": "Professionista abilitato (come da laurea in giurisprudenza o esame di avvocato) e abilitato all'esercizio della professione forense, ovvero rappresentare le parti in giudizio o in giudizio e prestare consulenza legale.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/lawyer-us.mp3"
			},
			{
				"id": "4",
				"english": "particularly",
				"foreignword": "in particolar modo",
				"examplesentence_english": "The apéritifs were particularly stimulating.",
				"examplesentence_foreign": "Gli aperitivi sono stati particolarmente stimolanti.",
				"english_definition": "(focus) Especially, extremely.",
				"foreign_definition": "(focus) Soprattutto, estremamente.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/particularly-us.mp3"
			},
			{
				"id": "5",
				"english": "gender",
				"foreignword": "genere",
				"examplesentence_english": "The effect of the medication is dependent upon age, gender, and other factors.",
				"examplesentence_foreign": "L'effetto del farmaco dipende dall'età, dal sesso e da altri fattori.",
				"english_definition": "Class; kind.",
				"foreign_definition": "Classe; Tipo.",
				"audiofilelink": ""
			},
			{
				"id": "6",
				"english": "literary",
				"foreignword": "letterario",
				"examplesentence_english": "a literary history",
				"examplesentence_foreign": "una storia letteraria",
				"english_definition": "Relating to literature.",
				"foreign_definition": "Relativo alla letteratura.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3"
			},
			{
				"id": "7",
				"english": "cotton",
				"foreignword": "cotone",
				"examplesentence_english": "",
				"examplesentence_foreign": "",
				"english_definition": "Gossypium, a genus of plant used as a source of cotton fiber.",
				"foreign_definition": "Gossypium, un genere di pianta utilizzata come fonte di fibra di cotone.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/cotton-1-us.mp3"
			},
			{
				"id": "8",
				"english": "station",
				"foreignword": "stazione",
				"examplesentence_english": "She had ambitions beyond her station.",
				"examplesentence_foreign": "Aveva ambizioni al di là del suo rango.",
				"english_definition": "A stopping place.",
				"foreign_definition": "Un luogo di sosta.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/station-au.mp3"
			},
			{
				"id": "9",
				"english": "everyone",
				"foreignword": "tutti",
				"examplesentence_english": "",
				"examplesentence_foreign": "",
				"english_definition": "Every person.",
				"foreign_definition": "Ogni persona.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/everyone-us.mp3"
			},
			{
				"id": "10",
				"english": "life",
				"foreignword": "vita",
				"examplesentence_english": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.",
				"examplesentence_foreign": "Avendo sperimentato entrambi, il vampiro decise che preferiva la (non)morte alla vita. Ha rinunciato alla vita.",
				"english_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.",
				"foreign_definition": "Lo stato degli organismi che precede la loro morte, caratterizzato da processi biologici come il metabolismo e la riproduzione e che li distingue dagli oggetti inanimati; lo stato di essere vivo e vivente.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/life-uk.mp3"
			}
		],
		"date": "03-29-2023"
	}`), &expected); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	var actual expectedTenWordResponse
	if err := json.Unmarshal([]byte(rr.Body.String()), &actual); err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %+v want %+v", actual, expected)
	}
}
