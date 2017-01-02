package simpleping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/stvp/rollbar"
)

func initializeGetStartedButton() error {
	getStarted := GetStartedTemplate{}

	getStarted.SettingType = "call_to_actions"
	getStarted.ThreadState = "new_thread"
	getStarted.CallToActions = []CTA{
		0: CTA{
			Payload: "subscribe_new_thread",
		},
	}

	getStartedB, err := json.Marshal(getStarted)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("FB_GRAPH_API_URL")+"/thread_settings", bytes.NewBuffer(getStartedB))
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

	bodyB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to initialize Get Started button for new threads: %s", string(bodyB))
		rollbar.Error(rollbar.ERR, err)
		return err
	}

	return nil
}
