package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handler() {

}

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
	gorillaRoot := mux.NewRouter()
	gorillaRoot.HandleFunc("api/user/{id:[0]}", Yazdir)
	http.Handle("/", gorillaRoot)

	http.ListenAndServe(":8888", nil)

}
