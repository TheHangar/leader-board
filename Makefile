build:
	@go build -o bin/leaderboard

run: build
	@./bin/leaderboard

test:
	@go test -v ./...
