# TenWords App

Use the app here! https://tenwords.netlify.app/words

Project Name: TenWords

Project Description:  TenWords is a language-learning web application that allows users to quickly learn a language with just 10 words a day. After creating an account, users will be able to select a language that would like to learn. The languages included are Spanish, French, Russian, Italian, Japanese, and Chinese. Each day, users will receive a new packet of 10 words for the day. On our website, users will be able to learn the 10 words in the language through flashcards and quizzes. The flashcards will include text to speech options, so users can hear the pronunication of the words.

Back-end team:
Grace Hu,
Aayesha Islam

Front-end team:
Lindsey Seay, 
Olivia Bronstein

## Requirements for Running and Testing Application
To run the application, clone the repository.
`cd` into the directory in which you've cloned the repository.
- Type `npm install` to install dependencies for the frontend project
- Type `npm run start` to run the frontend application 
- Run the commands below to install the Go packages
  `go get github.com/Conight/go-googletrans`
  `go get github.com/gen2brain/beeep`
  `go get github.com/gorilla/mux`
  `go get go.mongodb.org/mongo-driver/bson`
  `go get go.mongodb.org/mongo-driver/mongo`
  `go get go.mongodb.org/mongo-driver/mongo/options`
- Type  `go build` and then `./introtoSWE` to run the backend application

Go to `http://localhost:3000` to interact with the application.
