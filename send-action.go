package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/stvp/rollbar"
)

func sendAction(w http.ResponseWriter, recipientID string, action string) error {
	var err error
	if recipientID == "" {
		err = fmt.Errorf("recipientId is required")
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	if action == "" {
		err = fmt.Errorf("action is required")
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	outgoingMessage := OutgoingMessage{}
	outgoingMessage.Recipient.ID = recipientID
	outgoingMessage.SenderAction = action

	bodyB, err := json.Marshal(outgoingMessage)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("FB_MESSENGER_URL"), bytes.NewBuffer(bodyB))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	params := url.Values{}
	params.Set("access_token", os.Getenv("FB_PAGE_ACCESS_TOKEN"))

	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	defer resp.Body.Close()

	log.Printf("SEND ACTION BODY %v", string(bodyB))

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to update action: %s successfully - resp %v", action, resp)
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	return nil
}
