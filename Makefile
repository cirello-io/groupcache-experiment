all:
	go build ./cmd/backend
	go build ./cmd/frontend
	go build ./cmd/cli

client:
	echo 'set'
	./cli -set -k foo -v bar
	echo 'slow get'
	./cli -get -k foo
	echo 'fast get'
	./cli -get -k foo
