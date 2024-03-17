package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    }).Methods("GET")

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("PUT request received"))
    }).Methods("PUT")

    r.HandleFunc("/backend-api/conversation", func(w http.ResponseWriter, r *http.Request) {
        var reqBody struct {
            Message string `json:"message"`
        }
        if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        log.Println(reqBody.Message)
        json.NewEncoder(w).Encode(map[string]string{"msg": "2+2=4"})
    }).Methods("POST")

    log.Fatal(http.ListenAndServe(":3011", r))
}
