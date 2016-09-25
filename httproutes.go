package main

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type HTMLContent struct {
	CSS template.CSS
}

func registerRoutes(router *mux.Router, store *ForwardStore) {
	// Serve static files from public folder
	box := rice.MustFindBox("public")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(box.HTTPBox())))

	// Routes where the index files just needs to be show.
	// Angular will sort the logic on the front end.
	router.PathPrefix("/set").HandlerFunc(httpServeIndex)
	router.PathPrefix("/recent").HandlerFunc(httpServeIndex)

	// API router
	router.HandleFunc("/api/get/{uid}", httpGetForward(store)).Methods("Get")
	router.HandleFunc("/api/set", httpSetForward(store)).Methods("Post")
	router.HandleFunc("/api/recents", httpGetRecents(store)).Methods("Get")

	// Redirect on forwarder access
	router.HandleFunc("/{uid}", httpForward(store)).Methods("Get")
	router.HandleFunc("/", httpForward(store)).Methods("Get")
}

// Simple handler for serving the index.html file
func httpServeIndex(w http.ResponseWriter, r *http.Request) {
	box := rice.MustFindBox("public")
	html := box.MustString("index.html")
	css := box.MustString("dist/styles.css")
	t := template.New("index")
	t, _ = t.Parse(html)
	t.Execute(w, &HTMLContent{CSS: template.CSS(css)})
}

// Enpoint to perform the forward redirect
func httpForward(store *ForwardStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["uid"]
		if key == "" {
			key = "0"
		}
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

		newForward := store.set(f)
		json.NewEncoder(w).Encode(newForward)
	}
}
