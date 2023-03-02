# Sprint 2
# Frontend Branch Commits (Link: https://github.com/gracelhu/introtoSWE/tree/sprint-2-frontend)
## Frontend Unit Tests (https://github.com/gracelhu/introtoSWE/tree/sprint-2-frontend/src/tests)
  ### Unit Tests 
   - Words Page Component
      - Tests for Changing Language Selection Adjusts Subtitle: Ensure that changing the selected language updates the language displayed on the page
      - Tests for Changing Language Selection Changes Words Table Component: Ensure that changing the selected language updates the language displayed on the table
   - Words Table Component
      - Ensure exactly 10 words show up in table
      - Ensure words are in passed in language (Italian)
      - Ensure words are in passed in language (French)
  ### Cypress Tests
  - Sign-In Page
	- Test entering login and submitting
  - Sign-Up Page
	- Test entering login and submitting
  - Words Page
	- Test that connection has been made to API
  - Word Table
	- Test word translation matches up
	- Test the word audio button
  
### Backend Branck Link: https://github.com/gracelhu/TenWords/tree/sprint-2-backend

## Backend Unit Tests

- func TestWord: Tests Word struct and its json components, checks whether it is compatible with the MongoDB Database. Test Data was sent to the Test Database and it was retrieved and compared with the original data and the test passes.

- func TestUpdateWordProgress: ProgressIndex tracks the number of words completed by the User and the MapDatetoIndex matches the date with the ProgressIndex to keep track of dates so that users can check their previous days' progress. The test function, tests whether the ProgressIndex is updated in the database every time the api is called.

- func TestGetWord: Tests the route handler which sends a single word with all the required components.

- func TestGetTenWordsByID: Tests the route handler which sends ten words with all the required components.



## Documentation for Backend API: 

 TenWords REST API Documentation
Where does the TenWords REST API get the vocabulary words from?

In our project, there is a text file we created called “wordlist.txt” with about 3,000 commonly used English vocabulary words. Every word is indexed.
If an API call is trying to fetch information on a single word, the endpoint contains the index of the single word.
If an API call is trying to fetch information for a 10-word package, the endpoint contains the starting index of the package. For example, if I want the 10-word package with index 1, that means I want words 1-10.
TenWords REST API capabilities:

Can fetch single words from a list by their index in 6 different languages (Spanish, French, Russian, Italian, Japanese, and Chinese simplified). Each single word response is sent in JSON format. This is an example response: { "id": "1", "english": "abandon", "foreignword": "", "examplesentence_english": "Many baby girls have been abandoned on the streets of Beijing.", "examplesentence_foreign": "许多女婴被遗弃在北京街头。", "english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.", "foreign_definition": "放弃或放弃控制，投降或放弃自己，或屈服于自己的情绪。", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3" }

Can fetch 10-word packages from a list by their starting index in 6 different languages (Spanish, French, Russian, Italian, Japanese, and Chinese simplified). Records the date of the API call. Each 10-word package response is sent in JSON format. This is an example response: { "tenwords": [ { "id": "1", "english": "abandon", "foreignword": "abandonar", "examplesentence_english": "Many baby girls have been abandoned on the str eets of Beijing.", "examplesentence_foreign": "Muchas niñas han sido abandonadas en las calles de Beijing.", "english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.", "foreign_definition": "Renunciar o renunciar al control de, rendirse o entregarse, o ceder a las propias emociones.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3" }, { "id": "2", "english": "sudden", "foreignword": "repentino", "examplesentence_english": "The sudden drop in temperature left everyone cold and confused.", "examplesentence_foreign": "La repentina caída de la temperatura dejó a todos helados y confundidos.", "english_definition": "An unexpected occurrence; a surprise.", "foreign_definition": "Una ocurrencia inesperada; una sorpresa.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3" }, { "id": "3", "english": "lawyer", "foreignword": "abogado", "examplesentence_english": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation", "examplesentence_foreign": "El tiempo y el consejo de un abogado son su valor en el comercio. - aforismo a menudo acreditado a Abraham Lincoln, pero sin atestación", "english_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.", "foreign_definition": "Una persona profesional calificada (por un título en derecho o un examen de la barra) y autorizada para ejercer la abogacía, es decir, representar a las partes en demandas o juicios y brindar asesoramiento legal.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/lawyer-us.mp3" }, { "id": "4", "english": "particularly", "foreignword": "particularmente", "examplesentence_english": "The apéritifs were particularly stimulating.", "examplesentence_foreign": "Los aperitivos fueron particularmente estimulantes.", "english_definition": "(focus) Especially, extremely.", "foreign_definition": "(enfoque) Especialmente, extremadamente.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/particularly-us.mp3" }, { "id": "5", "english": "gender", "foreignword": "género", "examplesentence_english": "The effect of the medication is dependent upon age, gender, and other factors.", "examplesentence_foreign": "El efecto del medicamento depende de la edad, el sexo y otros factores.", "english_definition": "Class; kind.", "foreign_definition": "Clase; amable.", "audiofilelink": "" }, { "id": "6", "english": "literary", "foreignword": "literario", "examplesentence_english": "a literary history", "examplesentence_foreign": "una historia literaria", "english_definition": "Relating to literature.", "foreign_definition": "Relativo a la literatura.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3" }, { "id": "7", "english": "cotton", "foreignword": "algodón", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Gossypium, a genus of plant used as a source of cotton fiber.", "foreign_definition": "Gossypium, un género de planta utilizado como fuente de fibra de algodón.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/cotton-1-us.mp3" }, { "id": "8", "english": "station", "foreignword": "estación", "examplesentence_english": "She had ambitions beyond her station.", "examplesentence_foreign": "Ella tenía ambiciones más allá de su posición.", "english_definition": "A stopping place.", "foreign_definition": "Un lugar de parada.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/station-au.mp3" }, { "id": "9", "english": "everyone", "foreignword": "todos", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Every person.", "foreign_definition": "Cada persona.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/everyone-us.mp3" }, { "id": "10", "english": "life", "foreignword": "vida", "examplesentence_english": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.", "examplesentence_foreign": "Habiendo experimentado ambos, el vampiro decidió que prefería la (des)muerte a la vida. Renunció a la vida.", "english_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.", "foreign_definition": "El estado de los organismos que precede a su muerte, caracterizado por procesos biológicos como el metabolismo y la reproducción y que los distingue de los objetos inanimados; el estado de estar vivo y vivir.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/life-uk.mp3" } ], "date": "02-26-2023" }

Can fetch 10-word packages from a list by the date it was sent out. Follows the same JSON format response as the 10-word package response call by starting index.

TenWords REST API Route Handlers

There are 3 GET request route handlers in total

The route handlers will require you to pass in a parameter called {languagecode} which will determine which language to translate to. These are the language codes: • Spanish – es • French – fr • Russian – ru • Italian – it • Japanese – ja • Chinese - zh-cn

The 3 GET request route handlers are: • Getting a single word response in a language by the two parameters {languagecode} and {index}  Handled by the route handler function “getWord” • Getting a 10-word package response in a language by the two parameters {languagecode} and {index}  Handled by the route handler function “getTenWordsByID” • Getting a 10-word package response in a language by the two parameters {languagecode} and {date}  Handled by the route handler function “getTenWordsByDate” • The route handlers pass in a language code parameter to the route handler functions so it knows what language to translate to. An example of a route handler fetching a 10-word Spanish package: r.HandleFunc("/api/words/es/package/{id}", func(w http.ResponseWriter, r *http.Request) {getTenWordsByID(w, r, "es")}).Methods("GET")

How to use the TenWords REST API?

Get into the proper directory (restapi.go file), run “go build”, and run the generated executable (./executablename). The TenWords REST API is now running on port 8000. You can now make any necessary API calls.
To make an http GET request to fetch a 10-word package in one of the six language options, use the endpoint “http://localhost:8000/api/words/{languagecode}/package/{id}”
To make an http GET request to fetch a single word in one of the six language options, use the endpoint “http://localhost:8000/api/words/{languagecode}/single/{id}”
To make an http GET request to fetch a single word in one of the six language options, use the endpoint http://localhost:8000/api/words/{languagecode}/date/{date}

## Details of MongoDB Database

- The main thing we are tracking user's progress is by the ProgressIndex which signifies how many words they have completed. Everyday, the user gets a package of ten words and the index is updated by 10. 

- The Database stores 2 things: First is the ProgressIndex itself so that it is not reset everytime the api is called. Second is a Map which has a Key of Date and a value of ProgressIndex, so that later frontend can access the ProgressIndex for previous dates easily.

- In addition to inserting data, there is also retrieval functionality using the Find functionality. 

- Everytime, the api is called and the updated ProgressIndex and Map values are pushed to MongoDB and all the previous data in the collection is printed to the console.


<img width="416" alt="image" src="https://user-images.githubusercontent.com/92817486/222310526-adc79fb8-6ec1-419d-b4ba-3c749c8cd859.png">

