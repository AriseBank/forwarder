package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func registerRoutes(router *mux.Router, store *ForwardStore) {
	// Serve static files from public folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	router.PathPrefix("/set").HandlerFunc(httpServeIndex)

	// API router
	router.HandleFunc("/api/get/{uid}", httpGetForward(store)).Methods("Get")
	router.HandleFunc("/api/set/{uid}", httpSetForward(store)).Methods("Post")
	router.HandleFunc("/api/recents", httpGetRecents(store)).Methods("Get")

	// Redirect on forwarder access
	router.HandleFunc("/{uid}", httpForward(store)).Methods("Get")
}

// Simple handler for serving the index.html file
func httpServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

// Enpoint to perform the forward redirect
func httpForward(store *ForwardStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["uid"]
		forward, _ := store.get(key)
		http.Redirect(w, r, forward.URL, 302)
	}
}

// Endpoint for getting a single forward
func httpGetForward(store *ForwardStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["uid"]
		forward, _ := store.get(key)
		json.NewEncoder(w).Encode(forward)
	}
}

// Endpoint for getting the recently created forwards
func httpGetRecents(store *ForwardStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(store.Recents)
	}
}

// Endpoint for creating and setting forwards
func httpSetForward(store *ForwardStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["uid"]
		f := &Forward{}

		if r.Body == nil {
			http.Error(w, "No Post body", 400)
			return
		}

		err := json.NewDecoder(r.Body).Decode(f)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if key != f.UID {
			http.Error(w, "Incorrect Key Match", 500)
			return
		}

		store.set(f)
		json.NewEncoder(w).Encode(f)
	}
}
