package main

import (
	"fmt"
	"net/http"

	"github.com/stvp/rollbar"
)

func pingURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	for _, subscriber := range subcribers {
		message := Message{}
		message.Sender.ID = subscriber

		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("ping for %s returned %s :(", url, resp.Status)
			rollbar.Error(rollbar.ERR, err)
			message.MessageData.Text = err.Error()
			handleOutgoing(message)
		}
	}
}

func dailyUpdatePing(url string) {
	resp, err := http.Get(url)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	for _, subscriber := range subcribers {
		message := Message{}
		message.Sender.ID = subscriber

		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("daily update: ping for %s is unhealthy today - %s :(", url, resp.Status)
			rollbar.Error(rollbar.ERR, err)
			message.MessageData.Text = err.Error()
		}

		if resp.StatusCode == http.StatusOK {
			message.MessageData.Text = fmt.Sprintf("daily update: ping for %s is looking healthy today!", url)
		}

		handleOutgoing(message)
	}
}
