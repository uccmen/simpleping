package simpleping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/stvp/rollbar"
)

func handleOutgoing(message Message) {
	err := sendAction(message.Sender.ID, "typing_on")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}

	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	outgoingMessage := OutgoingMessage{}
	outgoingMessage.Recipient.ID = message.Sender.ID
	messageData := &OutgoingMessageData{}
	messageData.Text = message.MessageData.Text
	outgoingMessage.Message = messageData

	bodyB, err := json.Marshal(outgoingMessage)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	log.Println(string(bodyB))

	req, err := http.NewRequest("POST", os.Getenv("FB_GRAPH_API_URL")+"/messages", bytes.NewBuffer(bodyB))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	params := url.Values{}
	params.Set("access_token", os.Getenv("FB_PAGE_ACCESS_TOKEN"))

	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", "application/json")

	err = sendAction(message.Sender.ID, "typing_off")
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response failed to send successfully")
		rollbar.Error(rollbar.ERR, err)
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			rollbar.Error(rollbar.ERR, err)
			return
		}

		fmt.Printf("%q", dump)
		return
	}
}
