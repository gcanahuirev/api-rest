package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Shirt struct {
	Class     string `json:"class"`
	Material 	string `json:"material"`
	Id       	string `json:"id"`  
	Size     	int16  `json:"size"`
}

type shirtHandlers struct {
	sync.Mutex
	store map[string] Shirt
}
type adminPortal struct {
	password string
}

func (h *shirtHandlers) shirts(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
		case "GET":
			h.get(w,r)
			return
		case "POST":
			h.post(w, r)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("method not allowed"))
			return 
	}
}

func (h *shirtHandlers) get(w http.ResponseWriter, _ *http.Request)  {
	shirts := make([]Shirt, len(h.store))

	h.Lock()
	i := 0
	for _, shirt := range h.store {
		shirts[i] = shirt
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(shirts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *shirtHandlers) getRandomShirt(w http.ResponseWriter, _ *http.Request) {
	ids := make([]string, len(h.store))
	h.Lock()
	i := 0
	for id := range h.store {
		ids[i] = id
		i++
	}

	defer h.Unlock()

	var target string
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		target = ids[rand.Intn((len(ids)))]
	}

	w.Header().Add("location", fmt.Sprintf("/shirts/%s", target))
	w.WriteHeader(http.StatusFound)
}
func (h *shirtHandlers) getShirt(w http.ResponseWriter, r *http.Request)  {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if parts[2] == "random"{
		h.getRandomShirt(w, r)
		return
	}

	h.Lock()
	shirt, ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(shirt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *shirtHandlers) post(w http.ResponseWriter, r *http.Request)  {
	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got'%s'", ct)))
		return
	}

	var shirt Shirt
	err = json.Unmarshal(bodyBytes, &shirt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	shirt.Id = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[shirt.Id] = shirt
	defer h.Unlock()
}

func newShirtHandlers() *shirtHandlers  {
	return &shirtHandlers{
		store: map[string]Shirt{},
	}
}
func newAdminPortal()  *adminPortal{
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("required env var ADMIN_PASSWORD not set")
	}
	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request){
	user, pass, ok := r.BasicAuth()

	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Write([]byte("<html><h1>secret admin portal</h1></html>"))
}

func main()  {
	admin := newAdminPortal()
	shirtHandlers := newShirtHandlers()
	http.HandleFunc("/shirts", shirtHandlers.shirts)
	http.HandleFunc("/shirts/", shirtHandlers.getShirt)
	http.HandleFunc("/admin", admin.handler)

	addr := "localhost:3000"
	log.Println("Server is running at", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}