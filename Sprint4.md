# Sprint 4
Link to Sprint 4 Branch: [https://github.com/gracelhu/TenWords/blob/sprint-4](https://github.com/gracelhu/TenWords/tree/sprint-4)

## Overview of Sprint Progress
This Sprint, we made significant progress on the application by adding new features that make the app better! One of the features added to the application is the highlight/underline feature. With this feature, users to clearly see the vocabulary words underlines and bolded in the context of example sentences. A calendar selection feature was also implemented that corresponds to user progress on word packages. This allows users to find word packages generated for them on specific dates. Additionally, now users will only see sign in/sign up options in the navigation menu when they are not currently signed in.

## Frontend
### Frontend Unit Tests
Link to Tests: https://github.com/gracelhu/TenWords/tree/sprint-4/src/tests
#### Words Page Component
  - Tests for Changing Language Selection Adjusts Subtitle: Ensure that changing the selected language updates the language displayed on the page
  - Tests for Changing Language Selection Changes Words Table Component: Ensure that changing the selected language updates the language displayed on the table
  - Tests for Calendar Feature Instructions
#### Words Table Component
  - Ensure exactly 10 words show up in table
  - Ensure words are in passed in language (Italian)
  - Ensure words are in passed in language (French)
#### Practice/Quiz Feature Component
  - Ensure that Quiz feature renders correctly by ensuring that quiz title/heading displays with accurate text reflecting topic of learning for user
  - Ensure that instructions for taking quiz/completing practice are shown to the user and enables them to understand what to do on Quiz page for all language selections
#### Navigation Menu
  - Ensure Home and Your Words Tabs correctly display

#### Cypress Unit Tests
#### Word Table Component
  - Test that word table displays properly and elements work
  - Test that words are fetched within time
#### Sign-In Component
  - Test that sign-in page elements work
  - Test that sign-up page elements work
#### Sign-Out Component
  - Test that signing out redirects to the Words page
#### Authentication Component (e2e)
  - Test that signing in with a stored username / password logs the user in properly and redirects
  - Test that attempting to sign in with an invalid username / password combination does not log the user in and does not redirect

## Backend
### Backend Unit Tests
Link to Tests: https://github.com/gracelhu/TenWords/blob/sprint-3/finaltest_test.go
  - func TestUsername: Tests whether username is correctly updated in database.
  - func TestPassword: Tests whether the password is correctly updated in database.
  - func TestWord: Tests Word struct and its json components, checks whether it is compatible with the Database. Test Data was sent to the Test Database and it was retrieved and compared with the original data and the test passes.

- func TestUpdateWordProgress: ProgressIndex tracks the number of words completed by the User and the MapDatetoIndex matches the date with the ProgressIndex to keep track of dates so that users can check their previous days' progress. The test function, tests whether the ProgressIndex is updated in the database every time the api is called.

- func TestGetChineseWord: Tests the route handler which sends a single word with all the required components. Uses Chinese as the language code.

- func TestGetSpanishWord: Tests the route handler which sends a single word with all the required components. Uses Spanish as the language code.

- func TestGetTenWordsByID: Tests the route handler which sends ten words with all the required components. Uses Italian as the language code. 



### Backend Documentation (Updated)

#### REST API Documentation
**Where does the TenWords REST API get the vocabulary words from?**

In our project, there is a text file we created called “wordlist.txt” with about 3,000 commonly used English vocabulary words. Every word is indexed.
If an API call is trying to fetch information on a single word, the endpoint contains the index of the single word.
If an API call is trying to fetch information for a 10-word package, the endpoint contains the starting index of the package. For example, if I want the 10-word package with index 1, that means I want words 1-10.
TenWords REST API capabilities:

- Can fetch single words from a list by their index in 6 different languages (Spanish, French, Russian, Italian, Japanese, and Chinese simplified). Each single word response is sent in JSON format. This is an example response: { "id": "1", "english": "abandon", "foreignword": "", "examplesentence_english": "Many baby girls have been abandoned on the streets of Beijing.", "examplesentence_foreign": "许多女婴被遗弃在北京街头。", "english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.", "foreign_definition": "放弃或放弃控制，投降或放弃自己，或屈服于自己的情绪。", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3" }

- Can fetch 10-word packages from a list by their starting index in 6 different languages (Spanish, French, Russian, Italian, Japanese, and Chinese simplified). Records the date of the API call. Each 10-word package response is sent in JSON format. This is an example response: { "tenwords": [ { "id": "1", "english": "abandon", "foreignword": "abandonar", "examplesentence_english": "Many baby girls have been abandoned on the str eets of Beijing.", "examplesentence_foreign": "Muchas niñas han sido abandonadas en las calles de Beijing.", "english_definition": "To give up or relinquish control of, to surrender or to give oneself over, or to yield to one's emotions.", "foreign_definition": "Renunciar o renunciar al control de, rendirse o entregarse, o ceder a las propias emociones.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/abandon-us.mp3" }, { "id": "2", "english": "sudden", "foreignword": "repentino", "examplesentence_english": "The sudden drop in temperature left everyone cold and confused.", "examplesentence_foreign": "La repentina caída de la temperatura dejó a todos helados y confundidos.", "english_definition": "An unexpected occurrence; a surprise.", "foreign_definition": "Una ocurrencia inesperada; una sorpresa.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/sudden-us.mp3" }, { "id": "3", "english": "lawyer", "foreignword": "abogado", "examplesentence_english": "A lawyer's time and advice are his stock in trade. - aphorism often credited to Abraham Lincoln, but without attestation", "examplesentence_foreign": "El tiempo y el consejo de un abogado son su valor en el comercio. - aforismo a menudo acreditado a Abraham Lincoln, pero sin atestación", "english_definition": "A professional person qualified (as by a law degree or bar exam) and authorized to practice law, i.e. represent parties in lawsuits or trials and give legal advice.", "foreign_definition": "Una persona profesional calificada (por un título en derecho o un examen de la barra) y autorizada para ejercer la abogacía, es decir, representar a las partes en demandas o juicios y brindar asesoramiento legal.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/lawyer-us.mp3" }, { "id": "4", "english": "particularly", "foreignword": "particularmente", "examplesentence_english": "The apéritifs were particularly stimulating.", "examplesentence_foreign": "Los aperitivos fueron particularmente estimulantes.", "english_definition": "(focus) Especially, extremely.", "foreign_definition": "(enfoque) Especialmente, extremadamente.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/particularly-us.mp3" }, { "id": "5", "english": "gender", "foreignword": "género", "examplesentence_english": "The effect of the medication is dependent upon age, gender, and other factors.", "examplesentence_foreign": "El efecto del medicamento depende de la edad, el sexo y otros factores.", "english_definition": "Class; kind.", "foreign_definition": "Clase; amable.", "audiofilelink": "" }, { "id": "6", "english": "literary", "foreignword": "literario", "examplesentence_english": "a literary history", "examplesentence_foreign": "una historia literaria", "english_definition": "Relating to literature.", "foreign_definition": "Relativo a la literatura.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/literary-us.mp3" }, { "id": "7", "english": "cotton", "foreignword": "algodón", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Gossypium, a genus of plant used as a source of cotton fiber.", "foreign_definition": "Gossypium, un género de planta utilizado como fuente de fibra de algodón.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/cotton-1-us.mp3" }, { "id": "8", "english": "station", "foreignword": "estación", "examplesentence_english": "She had ambitions beyond her station.", "examplesentence_foreign": "Ella tenía ambiciones más allá de su posición.", "english_definition": "A stopping place.", "foreign_definition": "Un lugar de parada.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/station-au.mp3" }, { "id": "9", "english": "everyone", "foreignword": "todos", "examplesentence_english": "", "examplesentence_foreign": "", "english_definition": "Every person.", "foreign_definition": "Cada persona.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/everyone-us.mp3" }, { "id": "10", "english": "life", "foreignword": "vida", "examplesentence_english": "Having experienced both, the vampire decided that he preferred (un)death to life.  He gave up on life.", "examplesentence_foreign": "Habiendo experimentado ambos, el vampiro decidió que prefería la (des)muerte a la vida. Renunció a la vida.", "english_definition": "The state of organisms preceding their death, characterized by biological processes such as metabolism and reproduction and distinguishing them from inanimate objects; the state of being alive and living.", "foreign_definition": "El estado de los organismos que precede a su muerte, caracterizado por procesos biológicos como el metabolismo y la reproducción y que los distingue de los objetos inanimados; el estado de estar vivo y vivir.", "audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/life-uk.mp3" } ], "date": "02-26-2023" }

- Can fetch 10-word packages from a list by the date it was sent out. Follows the same JSON format response as the 10-word package response call by starting index.

**TenWords REST API Route Handlers**

There are 3 GET request route handlers in total

The route handlers will require you to pass in a parameter called {languagecode} which will determine which language to translate to. These are the language codes: 
- Spanish – es
- French – fr
- Russian – ru
- Italian – it
- Japanese – ja
- Chinese - zh-cn

The 3 GET request route handlers are: • Getting a single word response in a language by the two parameters {languagecode} and {index}  Handled by the route handler function “getWord” • Getting a 10-word package response in a language by the two parameters {languagecode} and {index}  Handled by the route handler function “getTenWordsByID” • Getting a 10-word package response in a language by the two parameters {languagecode} and {date}  Handled by the route handler function “getTenWordsByDate” • The route handlers pass in a language code parameter to the route handler functions so it knows what language to translate to. An example of a route handler fetching a 10-word Spanish package: r.HandleFunc("/api/words/es/package/{id}", func(w http.ResponseWriter, r *http.Request) {getTenWordsByID(w, r, "es")}).Methods("GET")

Newly added route handler called "getnameandpass" for user authentication. This gets the username and password inputted by the user in the signup or sign in page.
Example: r.HandleFunc("/api/words/package/auth", getnameandpass).Methods("GET")

**How to use the TenWords REST API?**

Get into the proper directory (restapi.go file), run “go build”, and run the generated executable (./executablename). The TenWords REST API is now running on port 8000. You can now make any necessary API calls.

- To make an http GET request to fetch a 10-word package in one of the six language options, use the endpoint “http://localhost:8000/api/words/{languagecode}/package/{id}”

- To make an http GET request to fetch a single word in one of the six language options, use the endpoint “http://localhost:8000/api/words/{languagecode}/single/{id}”

- To make an http GET request to fetch a single word in one of the six language options, use the endpoint http://localhost:8000/api/words/{languagecode}/date/{date}

- To make an http GET request to fetch the username and password, use the endpoint "http://localhost:8000/api/words/package/auth"


#### Database Details
##### How User is Authenticated
  - The data stored was the user's credentials, including their username and password. The date they are trying to login is also stored, along with a map which maps from username to password.
  - A validation state is determined depending on whether the user is new, returning, or invalid. To determine this, the database is searched for the exact match of the username to the password and if either is incorrect, then the state is set to invalid, if it is correct and found, then the state is set to "returning", and if it is new, then the state is "new".
  - The state is fetched by the frontend using a routehandler.


##### How User Progress Data Is Stored
- The main thing we are tracking user's progress is by the ProgressIndex which signifies how many words they have completed. The user gets a package of ten words and the index is updated by 10. 

- The Database stores 2 things: First is the ProgressIndex itself so that it is not reset everytime the api is called. Second is a Map which has a Key of Date and a value of ProgressIndex, so that later frontend can access the ProgressIndex for previous dates easily.

- In addition to inserting data, there is also retrieval functionality using the Find functionality. 

- Everytime, the api is called and the updated ProgressIndex and Map values are pushed to database and all the previous data in the collection is printed to the console.


<img width="416" alt="image" src="https://user-images.githubusercontent.com/92817486/222310526-adc79fb8-6ec1-419d-b4ba-3c749c8cd859.png">


