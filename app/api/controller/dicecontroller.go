package controller

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/comethale/dice-online/app/api/shared/database"
	"github.com/comethale/dice-online/app/api/shared/repositories/usermanagement/repository"
)

// Game represents an instance of a game
type Game struct {
	Result        bool  `json:"Result"`
	Goal          int64 `json:"Goal"`
	Score         int64 `json:"Score"`
	UserID        int64 `json:"UserID"`
	UserHighScore int64 `json:"UserHighScore"`
}

// RollDice takes a GET request and returns a dice roll result; or takes a POST request and initiates a game
func RollDice(w http.ResponseWriter, r *http.Request) {
	var game Game

	switch r.Method {
	case http.MethodPost:
		rollDicePost(w, r, &game)
	case http.MethodGet:
		rollDiceGet(w, r, &game)
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// *****************************************************************************
// Helper Functions
// *****************************************************************************

// performs the POST actions for RollDice
func rollDicePost(w http.ResponseWriter, r *http.Request, game *Game) {
	repo := repository.NewUserRepository(database.POSTGRESQL)

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("userid"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game.UserID = int64(userID)

	goal, err := strconv.Atoi(r.FormValue("goal"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game.Goal = int64(goal)

	user, err := repo.Get(game.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game.UserHighScore = user.HighScore

	json, err := json.Marshal(game)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}

// performs the GET actions for RollDice
func rollDiceGet(w http.ResponseWriter, r *http.Request, game *Game) {
	var roll int64
	repo := repository.NewUserRepository(database.POSTGRESQL)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&game)

	if err != nil {
		panic(err)
	}

	roll = rand.Int63n(5) + 1 // random number between 1 and 6, inclusive

	if game.Goal == roll { // User won, increase the score
		game.Result = true
		game.Score++
	} else { // User lost, set the User's highest score

		highscore := int64(math.Max(float64(game.Score), float64(game.UserHighScore)))
		_, err := repo.Update("", "", "", highscore, game.UserID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json, err := json.Marshal(game)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}
