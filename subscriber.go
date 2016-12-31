package main

import (
	"fmt"
	"os"

	"github.com/stvp/rollbar"
)

func addSubscriber(subscriberID string) error {
	conn := redisInstance.DB().Get()
	defer conn.Close()
	_, err := conn.Do("SADD", "subscribers", subscriberID)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	return nil
}

func getSubscribers() ([]string, error) {
	conn := redisInstance.DB().Get()
	defer conn.Close()
	res, err := conn.Do("SMEMBERS", "subscribers")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return []string{}, err
	}
	reply := res.([]interface{})
	for _, data := range reply {
		subcribers = append(subcribers, string(data.([]byte)))
	}

	return subcribers, nil
}

func subscribeNewThread(message Message) {
	err := addSubscriber(message.Sender.ID)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}

	message.MessageData.Text = fmt.Sprintf("You are now subcribed to receive server status for %s", os.Getenv("URL_TO_PING"))
	handleOutgoing(message)
	return
}
