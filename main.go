package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

const (
	DB_IP = "35.196.170.187:3306"
	DB_NAME = "octohack"
	DB_USER = "bitly"
	DB_PASSWORD = "yPlkKtGqviLmz1yy"
)

var db *sql.DB

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func main() {
	openDB()
	defer db.Close()

	serve()
}

func openDB() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASSWORD, DB_IP, DB_NAME))
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
}

func serve() {
	router := httprouter.New()
	router.GET("/users/:username", GetUser)
	router.POST("/users", CreateUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("username")

	var ret User
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", name).Scan(&ret.Id, &ret.Username, &ret.Email)


	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, formatError("DB ERROR", err), 500)
			return
		}
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(ret)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "INVALID_JSON", 422)
		return
	}

	if u.Username == "" {
		http.Error(w, "INVALID_ARG_USERNAME", 400)
		return
	}

	if u.Email == "" {
		http.Error(w, "INVALID_ARG_EMAIL", 400)
		return
	}

	_, err = db.Exec(`INSERT INTO users (username, email) VALUES (?, ?)`, u.Username, u.Email)
	if err != nil {
		http.Error(w, formatError("DB ERROR", err), 500)
		return
	}

	w.Header().Set("Content-Type", "text; charset=UTF-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func formatError(msg string, err error) string {
	return fmt.Sprintf("%s: %s", msg, err.Error())
}