package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {

	portFlag := flag.Int("p", 8080, "Define the port to listen on")
	flag.Parse()

	// Create forward store and router
	store := NewForwardStore()
	router := mux.NewRouter()

	// Register all router routes
	registerRoutes(router, store)

	// Start watcher to clean up old forwards
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				store.cleanUp()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

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

	Life      int64 `json:"life"`
	CreatedAt int64 `json:"createdAt"`
	ExpiresAt int64 `json:"expiresAt"`
}

// NewForward created a new Forward based on the given UID and URL.
// A default URL is used if an empty string is provided.
func NewForward(UID string, URL string) *Forward {
	if URL == "" {
		URL = "http://www.example.com"
	}

	life := int64(3600) // One hour
	t := time.Now().UTC().Unix()

	return &Forward{
		UID:       UID,
		URL:       URL,
		Life:      life,
		CreatedAt: t,
		ExpiresAt: t + life,
	}
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
		forward.Life = 0
	}
	return forward, exists
}

// set sets a forward in the store
func (store *ForwardStore) set(f *Forward) *Forward {
	forward := NewForward(f.UID, f.URL)
	store.Forwards[f.UID] = forward
	store.addRecent(forward)
	return forward
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

// cleanUp removes expired forwards from the store
func (store *ForwardStore) cleanUp() {
	now := time.Now().UTC().Unix()
	newRecents := make([]*Forward, 0)

	// Remove from recents if existing
	for _, v := range store.Recents {
		if v.ExpiresAt > now {
			newRecents = append(newRecents, v)
		}
	}

	store.Recents = newRecents

	// Delete from map
	for k, v := range store.Forwards {
		if v.ExpiresAt < now {
			delete(store.Forwards, k)
		}
	}
}
