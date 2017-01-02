package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stvp/rollbar"
	sp "github.com/uccmen/simpleping"
)

func main() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/webhook", sp.FBWebhook)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Panicln("ListenAndServe: ", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("I'm OK!"))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
