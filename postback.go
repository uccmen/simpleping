package main

import (
	"fmt"
	"os"
)

func handlePostback(message Message) {
	switch message.Postback.Payload {
	case "subscribe_new_thread":
		subscribeNewThread(message)
		return
	default:
		//do nothing
	}
}

func subscribeNewThread(message Message) {
	subcribers = append(subcribers, message.Sender.ID)
	message.MessageData.Text = fmt.Sprintf("You are now subcribed to receive server status for %s", os.Getenv("URL_TO_PING"))
	handleOutgoing(message)
	return
}
