package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

type Structure struct {
	Data []byte
}

func main() {
	address := "127.0.0.1:7070"
	box := packr.NewBox("./assets")

	structure := &Structure{}

	data, err := box.Find("structure.json")
	if err != nil {
		log.Fatal(err)
	}

	structure.Data = data
	log.Printf("Structure content \n%s\n", structure.Data)

	staticServer := http.StripPrefix("assets", http.FileServer(box))

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticServer))

	log.Printf("Server started on %s\n", address)
	http.ListenAndServe(address, router)
}

func ping(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"alive": "true"}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
