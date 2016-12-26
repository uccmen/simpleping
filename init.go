package main

import (
	"log"
	"os"

	"github.com/stvp/rollbar"
)

func init() {
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
	if os.Getenv("FB_MESSENGER_URL") == "" {
		log.Panicln("FB_MESSENGER_URL not set")
	}
	rollbar.Token = os.Getenv("ROLLBAR_TOKEN")
	rollbar.Environment = os.Getenv("RELEASE_STAGE") // defaults to "development"
}
