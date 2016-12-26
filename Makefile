GO_BUILD_ENV := GOOS=linux GOARCH=amd64

default: buildmac

deps:
	go get github.com/bugsnag/bugsnag-go

bin/simpleping: *.go
	$(GO_BUILD_ENV) go build -v -o $@ $^

bin/simpleping-mac: *.go
	go build -race -v -o $@ $^

build: bin/simpleping

buildmac: bin/simpleping-mac

runnotify: buildmac
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

heroku: bin/simpleping
	heroku container:push web
