package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

var userCounter = 0

var cm sync.RWMutex

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hellohandler)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Listening on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("err occurred:", err.Error())
	}
}

// Handles requests to the home page.
func hellohandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "i hope i get a good job 😣😣 ")
}

// Creates and stores a new user.
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	cm.Lock()
	userCache[userCounter] = user
	userCounter++
	cm.Unlock()

	w.WriteHeader(http.StatusOK)
}

// Retrieves a user using their ID.
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cm.RLock()
	user, ok := userCache[id]
	cm.RUnlock()

	if !ok {
		http.Error(w, "My Nigga User Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Deletes a user using their ID.
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "My Nigga User Not Found", http.StatusNotFound)
		return
	}

	cm.Lock()
	delete(userCache, id)
	cm.Unlock()

	w.WriteHeader(http.StatusOK)
}
