package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	Message string `json:"message"`
}

func Yazdir(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "kullanici ID: " + id

	message := API{messageConcat}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("hata oluştu")
	}
	fmt.Fprintf(w, string(output))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/{id}", Yazdir)
	r.HandleFunc("/lokum", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "adana adana merkez")
	})
	http.Handle("/", r)

	http.ListenAndServe("127.0.0.1:9090", r)

}
