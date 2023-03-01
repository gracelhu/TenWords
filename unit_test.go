package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func TestUpdateWordProgress(t *testing.T) {
	// Connect to the test database
	def := "testindex"


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
"
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

func TestGetWord(t *testing.T) {
	//var word Word
	// Set up a mock request and response
	
	req, err := http.NewRequest("GET", "/api/words/en/single/22", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a test router and call the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/single/{id}", getWord)

	router.ServeHTTP(rr, req)

	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
		"id": "22",
		"english": "surface",
		"foreignword": "surface",
		"examplesentence_english": "On the surface, the spy looked like a typical businessman.",
		"examplesentence_foreign": "On the surface, the spy looked like a typical businessman.",
		"english_definition": "The overside or up-side of a flat object such as a table, or of a liquid.",
		"foreign_definition": "The overside or up-side of a flat object such as a table, or of a liquid.",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/surface-us.mp3"
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		fmt.Println("body: ", rr.Body.String())
	}

}


func TestGetTenWordsByID(t *testing.T) {
	// Set up a mock request and response
	req, err := http.NewRequest("GET", "/api/words/en/package/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a test router and call the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/package/{id}", getTenWordsByID)

	router.ServeHTTP(rr, req)

	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
		"tenwords": [
			{
				"id": "1",
				"english": "abandon",
				"foreignword": "abandon",
				"examplesentence_english": "Many baby girls have been abandoned on the streets of Beijing.",
				"examplesentence_foreign": "Many baby girls have been abandoned on the streets of Beijing.",
				"english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.",
				"foreign_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3"
			},
			{
				"id": "2",
				"english": "sudden",
				"foreignword": "sudden",
				"examplesentence_english": "The sudden drop in temperature left everyone cold and confused.",
				"examplesentence_foreign": "The sudden drop in temperature left everyone cold and confused.",
				"english_definition": "An unexpected occurrence; a surprise.",
				"foreign_definition": "An unexpected occurrence; a surprise.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3"
			},
			{
				"id": "3",
				"english": "lawyer",
				"foreignword": "lawyer",
				"examplesentence_english": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation",
				"examplesentence_foreign": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation",
				"english_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.",
				"foreign_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/lawyer-us.mp3"
			},
			{
				"id": "4",
				"english": "particularly",
				"foreignword": "particularly",
				"examplesentence_english": "The apéritifs were particularly stimulating.",
				"examplesentence_foreign": "The apéritifs were particularly stimulating.",
				"english_definition": "(focus) Especially, extremely.",
				"foreign_definition": "(focus) Especially, extremely.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/particularly-us.mp3"
			},
			{
				"id": "5",
				"english": "gender",
				"foreignword": "gender",
				"examplesentence_english": "The effect of the medication is dependent upon age, gender, and other factors.",
				"examplesentence_foreign": "The effect of the medication is dependent upon age, gender, and other factors.",
				"english_definition": "Class; kind.",
				"foreign_definition": "Class; kind.",
				"audiofilelink": ""
			},
			{
				"id": "6",
				"english": "literary",
				"foreignword": "literary",
				"examplesentence_english": "a literary history",
				"examplesentence_foreign": "a literary history",
				"english_definition": "Relating to literature.",
				"foreign_definition": "Relating to literature.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3"
			},
			{
				"id": "7",
				"english": "cotton",
				"foreignword": "cotton",
				"examplesentence_english": "",
				"examplesentence_foreign": "",
				"english_definition": "Gossypium, a genus of plant used as a source of cotton fiber.",
				"foreign_definition": "Gossypium, a genus of plant used as a source of cotton fiber.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/cotton-1-us.mp3"
			},
			{
				"id": "8",
				"english": "station",
				"foreignword": "station",
				"examplesentence_english": "She had ambitions beyond her station.",
				"examplesentence_foreign": "She had ambitions beyond her station.",
				"english_definition": "A stopping place.",
				"foreign_definition": "A stopping place.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/station-au.mp3"
			},
			{
				"id": "9",
				"english": "everyone",
				"foreignword": "everyone",
				"examplesentence_english": "",
				"examplesentence_foreign": "",
				"english_definition": "Every person.",
				"foreign_definition": "Every person.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/everyone-us.mp3"
			},
			{
				"id": "10",
				"english": "life",
				"foreignword": "life",
				"examplesentence_english": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.",
				"examplesentence_foreign": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.",
				"english_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.",
				"foreign_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/life-uk.mp3"
			}
		],
		"date": "02-28-2023"
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		fmt.Println("body: ", rr.Body.String())
	}

}
