package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (h *shirtHandlers) get(w http.ResponseWriter, r *http.Request)  {
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

func (h *shirtHandlers) post(w http.ResponseWriter, r *http.Request)  {
	bodyBytes, err := ioutil.ReadAll(r.Body)
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
		store: map[string]Shirt{
			"id1": {
				Class   : "Manga Larga",
				Material: "Lana",
				Id      : "0001",
				Size    : 14,
			},
		},
	}
}

func main()  {
	shirtHandlers := newShirtHandlers()
	http.HandleFunc("/shirts", shirtHandlers.shirts)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		panic(err)
	}
}