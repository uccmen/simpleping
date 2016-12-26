package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stvp/rollbar"
)

func pingUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("ping for %s returned %s :(", url, resp.Status)
		for _, subscriber := range subcribers {
			message := Message{}
			message.MessageData.Text = err.Error()
			message.Sender.ID = subscriber
			handleOutgoing(message)
		}
		rollbar.Error(rollbar.ERR, err)
		return
	}

	//DEBUG
	log.Printf("SUBSCRIBERS: %v", subcribers)
	log.Println(fmt.Sprintf("ping for %s returned %d - %s", url, resp.StatusCode, resp.Status))
}
