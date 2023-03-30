package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	translator "github.com/Conight/go-googletrans"
	"github.com/gorilla/mux"
)

/* func TestUpdateWordProgress(t *testing.T) {
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
	//router.HandleFunc("/api/words/{languagecode}/single/{id}", getWord).Methods("GET")
	router.ServeHTTP(rr, req)
	//handler := http.HandlerFunc(getTenWordsByID)

	//handler.ServeHTTP(rr, req)
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
	//log.Fatal(http.ListenAndServe(":8000", router))
}
*/

/* func TestGetWord(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/words/zh/single/11", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getWord)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{
		"id": "11",
		"english": "reveal",
		"foreignword": "揭示",
		"examplesentence_english": "The comedian had been telling us about his sleep being disturbed by noise. Then came the reveal: he was sleeping on a bed in a department store.",
		"examplesentence_foreign": "喜剧演员一直告诉我们他的睡眠被噪音打扰了。然后揭露：他睡在一家百货公司的床上。",
		"english_definition": "The outer side of a window or door frame; the jamb.",
		"foreign_definition": "窗户或门框的外侧；门框。",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/reveal-au.mp3"
	}`

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &bodyMap); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	var expectedMap map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedMap); err != nil {
		t.Errorf("failed to unmarshal expected response body: %v", err)
	}

	if !reflect.DeepEqual(bodyMap, expectedMap) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			bodyMap, expectedMap)
	}
} */

/*func TestGetWord(t *testing.T) {

	//pass this small slice (first ten words in allWords) into the request
	mockWords := []Word{
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

} */

func TestGetWord(t *testing.T) {
	// Define the mockWords slice
	mockWords := []Word{
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

	expected := `{"id":"2","english":"sudden","foreignword":"突然的","examplesentence_english":"The sudden drop in temperature left everyone cold and confused.","examplesentence_foreign":"突如其来的降温，让所有人都感到寒冷和迷茫。","english_definition":"An unexpected occurrence; a surprise.","foreign_definition":"意外事件；惊喜。","audiofilelink":"https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		fmt.Println("body: ", rr.Body.String())
	}
}
