package main

import (
	"os"

	sp "github.com/uccmen/simpleping"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "daily" {
		sp.DailyUpdatePing(os.Getenv("URL_TO_PING"))
	}
}
