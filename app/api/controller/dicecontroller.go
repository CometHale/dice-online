package controller

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/comethale/dice-online/app/api/shared/database"
	"github.com/comethale/dice-online/app/api/shared/repositories/usermanagement/repository"
	glog "google.golang.org/appengine/log"
)

// Game represents an instance of a game
type Game struct {
	Result        bool  `json:"Result"`
	Roll          int64 `json:"Roll"`
	Goal          int64 `json:"Goal"`
	Score         int64 `json:"Score"`
	UserID        int64 `json:"UserID"`
	UserHighScore int64 `json:"UserHighScore"`
}

// StartGame takes a POST request and initiates a game
func StartGame(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		repo := repository.NewUserRepository(database.POSTGRESQL)
		game, err := setUpGame(r, repo)

		if err != nil {
			log.Println(err)
			glog.Errorf(nil, err.Error(), nil)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(game)

		if err != nil {
			log.Println(err)
			glog.Errorf(nil, err.Error(), nil)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	default:
		log.Println("405 Method Not Allowed")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// RollDice takes a POST request and returns a dice roll result
func RollDice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		repo := repository.NewUserRepository(database.POSTGRESQL)
		game, err := setUpGame(r, repo)

		game.Roll = rand.Int63n(5) + 1 // random number between 1 and 6, inclusive

		if game.Goal == game.Roll { // User won, increase the score
			game.Result = true
			game.Score++
		} else { // User lost, set the User's highest score

			highscore := int64(math.Max(float64(game.Score), float64(game.UserHighScore)))
			_, err := repo.Update("", "", highscore, game.UserID)

			if err != nil {
				log.Println(err)
				glog.Errorf(nil, err.Error(), nil)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		json, err := json.Marshal(game)

		if err != nil {
			log.Println(err)
			glog.Errorf(nil, err.Error(), nil)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	default:
		log.Println("405 Method Not Allowed")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// *****************************************************************************
// Helper Functions
// *****************************************************************************

// setUpGame parses the game information from the request and stores it in a Game struct
func setUpGame(r *http.Request, repo *repository.UserRepository) (*Game, error) {
	var game Game

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		glog.Errorf(nil, err.Error(), nil)
		return nil, err
	}

	userID, err := strconv.Atoi(r.PostFormValue("userid"))

	if err != nil {
		log.Println(err)
		glog.Errorf(nil, err.Error(), nil)
		return nil, err
	}

	game.UserID = int64(userID)
	goal, err := strconv.Atoi(r.PostFormValue("goal"))

	if err != nil {
		log.Println(err)
		glog.Errorf(nil, err.Error(), nil)
		return nil, err
	}

	game.Goal = int64(goal)

	user, err := repo.Get(game.UserID)
	if err != nil {
		log.Println(err)
		glog.Errorf(nil, err.Error(), nil)
		return nil, err
	}

	game.UserHighScore = user.HighScore

	return &game, nil
}
