package main

/*
The purpose of this file is to provide the correct structs to "unmarshall" the json response data
from a free api called "freedictionaryapi" --> See this link:
https://dictionaryapi.dev/

This API will allow us to fetch a word's definition, example sentence, synonyms, antonyms, audio file of
its pronounciation, etc. These fields will be displayed on the flash cards in the application

NOTE: The JSON names in your struct MUST match the JSON names in the response
Also, it's required to capitalize variable names in a struct
*/

// I need to unmarshall the json response data
type Words struct {
	Word      string     `json:"word"`
	Phonetics []Phonetic `json:"phonetics"`
	Meanings  []Meaning  `json:"meanings"`
}

type Phonetic struct {
	Audio string `json:"audio"`
}

type Meaning struct {
	Definitions []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}
