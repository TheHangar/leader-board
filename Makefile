build:
	@go build -o bin/leaderboard

run: build
	@./bin/leaderboard

build-img: build
	@docker build -t ghcr.io/thehangar/leader-board:$(IMG_VERSION) .

test:
	@go test -v ./...
