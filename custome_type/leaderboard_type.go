package custometype

import "time"

type Game struct {
    ApiKey string
    UUID string
    Name string
}

type Leaderboard struct {
    ID int `json:"-"`
    Game_UUID string `json:"-"`
    User_UUID string `json:"pseudo"`
    Score float64 `json:"score"`
    CreatedAt time.Time `json:"-"`
    Rank int `json:"rank"`
}
