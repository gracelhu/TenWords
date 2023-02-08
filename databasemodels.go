package main

   

type Word struct {
	ID              string `json:"id"`
	Word            string `json:"word"`
	Spanish         string `json:"spanish"`
	Definition      string `json:"definition"`
	Partofspeech    string `json:"partofspeech"`
	Examplesentence string `json:"examplesentence"`
}

type TenWordPackage struct {
	Tenwords []Word `json:"tenwords"`
	Date     string `json:"date"` //in this format: 01-02-2006
}
