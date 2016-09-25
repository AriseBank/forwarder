package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	portFlag := flag.Int("p", 8080, "Define the port to listen on")
	flag.Parse()

	// Create forward store and router
	store := NewForwardStore()
	router := mux.NewRouter()

	// Register all router routes
	registerRoutes(router, store)

	// Start server
	fmt.Println(fmt.Sprintf("Listening on port %d", *portFlag))
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), router)
}

/**
 * Foward Components
 */

// Forward represents a single url forwarder and conains details
// of the key it's set up on and the location it points to.
type Forward struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

// NewForward created a new Forward based on the given UID and URL.
// A default URL is used if an empty string is provided.
func NewForward(UID string, URL string) *Forward {
	if URL == "" {
		URL = "http://www.example.com"
	}
	return &Forward{UID: UID, URL: URL}
}

/**
 * Foward Store Components
 */
func NewForwardStore() *ForwardStore {
	s := &ForwardStore{
		Forwards:   make(map[string]*Forward),
		Recents:    make([]*Forward, 0),
		MaxRecents: 25,
	}
	return s
}

type ForwardStore struct {
	Forwards   map[string]*Forward
	Recents    []*Forward
	MaxRecents int
}

// get returns a forward for the given UID key,
// A default Forward is returned if it does not currently exist.
func (store *ForwardStore) get(key string) (*Forward, bool) {
	forward, exists := store.Forwards[key]
	if !exists {
		forward = NewForward(key, "")
	}
	return forward, exists
}

// set sets a forward in the store
func (store *ForwardStore) set(f *Forward) {
	store.Forwards[f.UID] = f
	store.addRecent(f)
}

// Add a new recent item to the store.
// Manages the recents to ensure it does not overflow.
func (store *ForwardStore) addRecent(f *Forward) {
	exists, index := false, -1
	for i, forward := range store.Recents {
		if forward.UID == f.UID {
			exists, index = true, i
		}
	}

	if exists {
		store.Recents = append(store.Recents[:index], store.Recents[index+1:]...)
	}

	if len(store.Recents) >= store.MaxRecents {
		store.Recents = append([]*Forward{f}, store.Recents[:store.MaxRecents-1]...)
	} else {
		store.Recents = append([]*Forward{f}, store.Recents...)
	}
}
