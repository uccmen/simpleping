package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/robfig/cron"
	"github.com/stvp/rollbar"
)

var subcribers []string
var location *time.Location
var pingCron *cron.Cron

func init() {
	var err error
	if os.Getenv("ROLLBAR_TOKEN") == "" {
		log.Panicln("ROLLBAR_TOKEN not set")
	}
	if os.Getenv("PORT") == "" {
		log.Panicln("PORT not set")
	}
	if os.Getenv("HUB_VERIFY_TOKEN") == "" {
		log.Panicln("HUB_VERIFY_TOKEN not set")
	}
	if os.Getenv("FB_PAGE_ACCESS_TOKEN") == "" {
		log.Panicln("FB_PAGE_ACCESS_TOKEN not set")
	}
	if os.Getenv("FB_GRAPH_API_URL") == "" {
		log.Panicln("FB_GRAPH_API_URL not set")
	}
	if os.Getenv("TIMEZONE") == "" { //IANA Time Zone e.g. "America/New_York"
		log.Panicln("TIMEZONE not set")
	}
	if os.Getenv("PING_CRON_EXPRESSION") == "" {
		log.Panicln("PING_CRON_EXPRESSION not set")
	}
	if os.Getenv("URL_TO_PING") == "" {
		log.Panicln("URL_TO_PING not set")
	}

	_, err = url.Parse(os.Getenv("URL_TO_PING"))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		log.Panicln(err)
	}

	location, err = time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		log.Panicln(err)
	}
	rollbar.Token = os.Getenv("ROLLBAR_TOKEN")
	rollbar.Environment = os.Getenv("RELEASE_STAGE") // defaults to "development"

	err = initializeGetStartedButton()
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		return
	}

	go schedulePing()
}

func schedulePing() {
	pingCron = cron.New()
	pingCron.AddFunc(os.Getenv("PING_CRON_EXPRESSION"), func() { pingUrl(os.Getenv("URL_TO_PING")) })
	pingCron.Start()
}