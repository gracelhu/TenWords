- A functional CRUD RESTFUL API for TenWords application
- This API fetches ten word packages from a text file called "wordlist.txt" and uses
golang googletrans library to translate the words into 6 different foreign languages:
spanish, french, russian, italian, japanese, and chinese.
- Every time a ten word package is fetched, the date that package was fetched is recorded and
stored in our database
- Additionally, it stores the language learning progress of users in a database by recording
what index in wordlist.txt the user left off on.
- This API also autogenerates example sentences for each vocabulary word by calling a repository
of millions of english sentences stored in a database

