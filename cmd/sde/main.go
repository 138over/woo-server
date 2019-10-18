package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type server struct {
    Port string
}

func main() {
    svr := &server{Port: "7667"}

    router := mux.NewRouter()
    router.HandleFunc("/ping", ping).Methods("GET")

    log.Printf("Server started on %s\n", svr.Port)
    http.ListenAndServe(":7667", router)
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
