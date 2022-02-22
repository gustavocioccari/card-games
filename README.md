## :file_folder: The Project
This project is an API to deal with card games.

## :rocket: Technologies
|   Back-End   |
| :---:        |
| Go           |
| Echo         |
| PostgreSQL   |
| Docker       |

## :computer: Installation
Set the PORT you want to start the API in a .env file on the root folder
___
Then run
```bash
docker-compose up -d --build
```
To set up the PostgreSQL database
___
Run
```bash
make migrate
make seed
```
It will run the model migrations and seed database with the default values, suit and the 52 default cards
___
To start the API run
```bash
make deploy
```
___
You can reach the following endpoints

- `POST toggl-cards/v1/decks`
It will create a default unshuffled deck
It accepts the parameters cards and shuffled (eg.: `?cards=AC,JS,5C,10D&shuffled=true`). If you pass them the deck will be created according to the parameters.
___
- `GET toggl-cards/v1/decks/:id/open`
It will return the deck according to its ID.
___
- `PATCH toggl-cards/v1/decks/:id/draw`
It will return the drawn cards, remove it from the deck and decrease the remaining cards.

### TODO
- Add more tests
- Add error handling when passing cards that don't exist as parameter
