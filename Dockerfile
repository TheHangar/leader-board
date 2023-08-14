FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o leaderboard .
EXPOSE 3000
CMD [ "./leaderboard" ]

LABEL org.opencontainers.image.source=https://github.com/thehangar/leadear-board
LABEL org.opencontainers.image.licenses=MIT
