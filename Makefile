build:
	@go build -o bin/leaderboard

run: build
	@./bin/leaderboard

build-img: build
	@docker build -t the-hangar-leaderboard:$(IMG_VERSION) .

test:
	@go test -v ./...
