package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//var idtoindex = make(map[int]int) {
//idtoindex[9]=8, }
type Maptry struct {
	Data map[int]interface{}
}
//der := &Maptry{}

type MongoField struct {
	
	Username  string `json: "Username"`
	Password  string   `json: "Password"`
	ProgressIndex int   `json: "ProgressIndex"`
	Data map[int]interface{} `json: "map"`

}

func main() {
	//var idtoindex = make(map[int]int) 
	//idtoindex[9]=8
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("First_Database").Collection("First COllection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := MongoField{
		Username:  "aayeshamislam",
		Password:  "XXXXXX",
		ProgressIndex: 1,
		//Data: [3]=7,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

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
	//col:= &idtoindex
	s:= &Maptry{}
	s.Data[8]=7
	s.Data[6]=3
	erp, errs:= col.InsertOne(ctx,s)
	if errs != nil {
		fmt.Println("InsertONE Error:", errs)
		os.Exit(1)
	}else{
	fmt.Println("InsertOne() result type: ", reflect.TypeOf(erp))
		fmt.Println("InsertOne() api result type: ", erp)

		newsID := erp.InsertedID
		fmt.Println("InsertedOne(), newID", newsID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newsID))
}
}
