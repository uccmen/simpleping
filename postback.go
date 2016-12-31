package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stvp/rollbar"
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

func addSubscriber(subscriberID string) error {
	conn := redisInstance.DB().Get()
	defer conn.Close()
	_, err := conn.Do("SADD", "subscribers", subscriberID)
	if err != nil {
		log.Println(err.Error())
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	return nil
}

func getSubscribers() error {
	conn := redisInstance.DB().Get()
	defer conn.Close()
	res, err := conn.Do("SMEMBERS", "subscribers")
	if err != nil {
		log.Println(err.Error())
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	log.Println(res)

	return nil
}

func subscribeNewThread(message Message) {
	err := addSubscriber(message.Sender.ID)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}

	err = getSubscribers()
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}
	message.MessageData.Text = fmt.Sprintf("You are now subcribed to receive server status for %s", os.Getenv("URL_TO_PING"))
	handleOutgoing(message)
	return
}
