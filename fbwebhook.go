package main

import "net/http"

func fbWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		confirmSubscription(w, r)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	// callbacks
	handleIncoming(w, r)
	return
}
