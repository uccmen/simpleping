package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/stvp/rollbar"
)

func handleOutgoing(w http.ResponseWriter, message Message) {
	err := sendAction(w, message.Sender.ID, "typing_on")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}
	time.Sleep(3 * time.Second)
	outgoingMessage := OutgoingMessage{}
	outgoingMessage.Recipient.ID = message.Sender.ID
	messageData := &OutgoingMessageData{}
	messageData.Text = message.MessageData.Text
	outgoingMessage.Message = messageData

	bodyB, err := json.Marshal(outgoingMessage)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	log.Println(string(bodyB))

	req, err := http.NewRequest("POST", os.Getenv("FB_MESSENGER_URL"), bytes.NewBuffer(bodyB))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Set("access_token", os.Getenv("FB_PAGE_ACCESS_TOKEN"))

	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", "application/json")

	err = sendAction(w, message.Sender.ID, "typing_off")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response failed to send successfully")
		rollbar.Error(rollbar.ERR, err)
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			rollbar.Error(rollbar.ERR, err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		fmt.Printf("%q", dump)
		return
	}
	w.WriteHeader(http.StatusOK)
}
