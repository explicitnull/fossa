build:
	go build -o ./bin/fossa ./cmd/server

run: build
	./bin/fossa -config=config/config.yaml
