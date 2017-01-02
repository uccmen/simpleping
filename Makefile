GO_BUILD_ENV := GOOS=linux GOARCH=amd64

default: buildmac

deps:
	go get github.com/stvp/rollbar
	go get github.com/robfig/cron

bin/simpleping: *.go
	$(GO_BUILD_ENV) go build -race -v -o $@ $^

bin/scheduleping: cmd/scheduleping/*.go
	$(GO_BUILD_ENV) go build -race -v -o $@ $^

bin/simpleping-mac: *.go
	go build -race -v -o $@ $^

bin/scheduleping-mac: cmd/scheduleping/*.go
	go build -race -v -o $@ $^

build: bin/simpleping bin/scheduleping

buildmac: bin/simpleping-mac bin/scheduleping-mac

lint: *.go
	golint

runnotify: lint buildmac
	-killall simpleping-mac
	-terminal-notifier -title "simpleping" -message "Built and running!" -remove
	bin/simpleping-mac

watch:
	supervisor --no-restart-on exit -e go,html -i bin --exec make -- runnotify

clean:
	rm -f bin/*

test:
	go test -v .

run: build init
	bin/simpleping

heroku: bin/simpleping bin/scheduleping
	heroku container:push web

