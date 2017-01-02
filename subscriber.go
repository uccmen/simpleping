package main

import (
	"fmt"
	"os"

	"github.com/stvp/rollbar"
	sp "github.com/uccmen/simpleping"
)

func addSubscriber(subscriberID string) error {
	conn := sp.RedisInstance.DB().Get()
	defer conn.Close()
	_, err := conn.Do("SADD", "subscribers", subscriberID)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	return nil
}

func getSubscribers() ([]string, error) {
	conn := sp.RedisInstance.DB().Get()
	defer conn.Close()
	res, err := conn.Do("SMEMBERS", "subscribers")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return []string{}, err
	}
	rawData := res.([]interface{})
	for _, rawSubscriber := range rawData {
		sp.Subcribers = append(sp.Subcribers, string(rawSubscriber.([]byte)))
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
