package main

import (
	"net/http"
	"os"
	"strings"
)

func confirmSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	hubMode := strings.TrimSpace(r.FormValue("hub.mode"))
	hubVerifyToken := strings.TrimSpace(r.FormValue("hub.verify_token"))
	hubChallenge := strings.TrimSpace(r.FormValue("hub.challenge"))

	if hubMode != "subscribe" {
		http.Error(w, "", http.StatusExpectationFailed)
		return
	}

	if hubVerifyToken != os.Getenv("HUB_VERIFY_TOKEN") {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	_, err := w.Write([]byte(hubChallenge))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
