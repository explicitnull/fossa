build:
	go build -o ./bin/fossa .

run: build
	./bin/fossa 
