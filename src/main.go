package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type todo struct {
	Id string	`json:"id"`
	Content string	`json:"content"`
	Completed bool	`json:"completed"`
}

var todos = []todo {
	{Id: "1", Content: "Clean room", Completed: true},
	{Id: "2", Content: "Remove tail", Completed: false},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.Id = strconv.Itoa(rand.Intn(10000000000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range todos {
		if item.Id == params["id"] {
			todos = append(todos[:i], todos[i + 1:]...)
			var todo todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.Id = strconv.Itoa(rand.Intn(10000000000))
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for i, item := range todos {
		if item.Id == params["id"] {
			todos = append(todos[:i], todos[i + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	
	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}