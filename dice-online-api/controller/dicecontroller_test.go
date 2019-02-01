package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"dice-online-api/controller"
	"dice-online-api/shared/database"
)

func TestStartGame(t *testing.T) {

	// set up mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	database.POSTGRESQL = db

	// User
	userID := 1 // test user
	userHighScore := 50
	userEmail := "123@gmail.com"
	username := "test"

	// Game
	goal := 2

	selectStatement := "SELECT email, username, highscore FROM users WHERE"
	rows := mock.NewRows([]string{"email", "username", "highscore"}).AddRow(userEmail, username, userHighScore)
	mock.ExpectQuery(selectStatement).WithArgs(userID).WillReturnRows(rows)

	// implements http.ResponseWriter, passed into the handler
	w := httptest.NewRecorder()

	// fake request to pass into the handler
	form := url.Values{}
	form.Add("userid", strconv.Itoa(userID))
	form.Add("goal", strconv.Itoa(goal))
	r := httptest.NewRequest("POST", "/start-game/", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Form = form

	controller.StartGame(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected: %d, Got: %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()
	if len(body) <= 0 {
		t.Fatalf("Expected Non-Empty Response Body")
	}

	expectedGame := &controller.Game{Goal: int64(goal), UserID: int64(userID), UserHighScore: int64(userHighScore)}
	expectedBody, err := json.Marshal(expectedGame)

	if body != string(expectedBody) {
		t.Fatalf("Got Unexpected Body:\n[%s]", body)
	}

}

func TestStartGameIncorrectRequestMethod(t *testing.T) {
	// implements http.ResponseWriter, passed into the handler
	w := httptest.NewRecorder()

	// fake request to pass into the handler
	r := httptest.NewRequest("GET", "/start-game/", nil)

	controller.StartGame(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected: %d, Got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	body := w.Body.String()
	if len(body) <= 0 {
		t.Fatalf("Expected Non-Empty Response Body")
	}

	expectedBody := "405 Method Not Allowed\n"

	if body != expectedBody {
		t.Fatalf("Got Unexpected Body:\n[%s]", body)
	}
}

func TestRollDice(t *testing.T) {

	// mock the db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	database.POSTGRESQL = db

	// User
	userID := 1 // test user
	userHighScore := 50
	userEmail := "123@gmail.com"
	username := "test"

	// Game
	goal := 2

	selectStatement := "SELECT email, username, highscore FROM users WHERE"
	rows := mock.NewRows([]string{"email", "username", "highscore"}).AddRow(userEmail, username, userHighScore)
	mock.ExpectQuery(selectStatement).WithArgs(userID).WillReturnRows(rows)

	updateRows := mock.NewRows([]string{"email", "username", "highscore"}).AddRow("", "", userHighScore)
	updateStatement := `^UPDATE users SET`
	mock.ExpectQuery(updateStatement).WithArgs(nil, nil, userHighScore, userID).WillReturnRows(updateRows)
	w := httptest.NewRecorder()

	// fake request to pass into the handler
	form := url.Values{}
	form.Add("userid", strconv.Itoa(userID))
	form.Add("goal", strconv.Itoa(goal))
	r := httptest.NewRequest("POST", "/start-game/", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Form = form

	controller.RollDice(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected: %d, Got: %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()
	if len(body) <= 0 {
		t.Fatalf("Expected Non-Empty Response Body")
	}
}

func TestRollDiceIncorrectRequestMethod(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/start-game/", nil)

	controller.RollDice(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected: %d, Got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	body := w.Body.String()
	if len(body) <= 0 {
		t.Fatalf("Expected Non-Empty Response Body")
	}

	expectedBody := "405 Method Not Allowed\n"

	if body != expectedBody {
		t.Fatalf("Got Unexpected Body:\n[%s]", body)
	}

}
