package custometype

import "time"

type Game struct {
    UUID string
    Name string
}

type Leaderboard struct {
    ID int
    Game_UUID string
    User_UUID string
    Score float64
    CreatedAt time.Time
}
