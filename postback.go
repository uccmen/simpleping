package main

func handlePostback(message Message) {
	switch message.Postback.Payload {
	case "subscribe_new_thread":
		subscribeNewThread(message)
		return
	default:
		//do nothing
	}
}
