package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	// Seed data
	posts = append(posts, Post{ID: 1, Title: "Meu primeiro post", Body: "Conteúdo do primeiro post"})

	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedPost Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i].Title = updatedPost.Title
			posts[i].Body = updatedPost.Body
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
