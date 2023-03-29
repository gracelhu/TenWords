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
  - TestUsername: Tests whether username is correctly updated in the database.
  - TestPassword: Tests whether the password is correctly updated in the database.
### Updated Backend Documentation
