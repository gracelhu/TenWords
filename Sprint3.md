# Sprint 3

## Overview of Sprint Progress
This Sprint, one of the core features added was the Practice feature. This feature allows users to practice translating ten word packets from english to the language of their choosing. After typing in corresponding translations for the words, users receive a score on their quiz, and receive feedback on what was answered correctly (displays in green) and incorrectly (displays in red). 

Additionally, the frontend and backend connection was solidified this sprint. An authentication server was made with a route handler and a database which was connected to the frontend which allows users to sign up and sign in by storing their credentials. 

## Frontend
### Frontend Unit Tests
#### Words Page Component
  - Tests for Changing Language Selection Adjusts Subtitle: Ensure that changing the selected language updates the language displayed on the page
  - Tests for Changing Language Selection Changes Words Table Component: Ensure that changing the selected language updates the language displayed on the table
#### Words Table Component
  - Ensure exactly 10 words show up in table
  - Ensure words are in passed in language (Italian)
  - Ensure words are in passed in language (French)
#### Practice/Quiz Feature Component
  - Ensure that Quiz feature renders correctly by ensuring that quiz title/heading displays with accurate text reflecting topic of learning for user
  - Ensure that instructions for taking quiz/completing practice are shown to user and enables them to understand what to do on Quiz page for all language selections

#### Cypress Unit Tests


## Backend
## Backend Unit Tests
  - func TestUsername: Tests whether username is correctly updated in the database.
  - func TestPassword: Tests whether the password is correctly updated in the database.
  - func TestWord: Tests Word struct and its json components, checks whether it is compatible with the MongoDB Database. Test Data was sent to the Test Database and it was retrieved and compared with the original data and the test passes.

- func TestUpdateWordProgress: ProgressIndex tracks the number of words completed by the User and the MapDatetoIndex matches the date with the ProgressIndex to keep track of dates so that users can check their previous days' progress. The test function, tests whether the ProgressIndex is updated in the database every time the api is called.

- func TestGetWord: Tests the route handler which sends a single word with all the required components.

- func TestGetTenWordsByID: Tests the route handler which sends ten words with all the required components.



### Updated Backend Documentation

## Details of MongoDB Database
#### Authentication Server
  - The data stored was the User's credentials, including their username and password. The date they are trying to login is also stored, along with a map which maps from username to password.
  - A validation state is determined depending on whether the user is new, returning, or invalid. To determine this, the database is searched for the exact match of the username to the password and if either is incorrect, then the state is set to invalid, if it is correct and found, then the state is set to "returning", and if it is new, then the state is "new".
  - The state is fetched by the frontend using a routehandler.
#### User Progress Data
- The main thing we are tracking user's progress is by the ProgressIndex which signifies how many words they have completed. The user gets a package of ten words and the index is updated by 10. 

- The Database stores 2 things: First is the ProgressIndex itself so that it is not reset everytime the api is called. Second is a Map which has a Key of Date and a value of ProgressIndex, so that later frontend can access the ProgressIndex for previous dates easily.

- In addition to inserting data, there is also retrieval functionality using the Find functionality. 

- Everytime, the api is called and the updated ProgressIndex and Map values are pushed to MongoDB and all the previous data in the collection is printed to the console.


<img width="416" alt="image" src="https://user-images.githubusercontent.com/92817486/222310526-adc79fb8-6ec1-419d-b4ba-3c749c8cd859.png">


