package store

import custometype "github.com/TheHangar/leader-board/custome_type"

type PostgresLeaderboardStore struct {
}

func NewPostgresLeaderboardStore() *PostgresLeaderboardStore {
    return &PostgresLeaderboardStore{}
}

func (s *PostgresLeaderboardStore) AddLeaderboard(lb *custometype.Leaderboard) error {
    db, err := postgresConnect()

    if err != nil {
        return err
    }

    defer db.Close()

    _, err = db.Exec("INSERT INTO leaderboard (pseudo, score, game_uuid) VALUES ($1, $2, $3);", lb.User_UUID, lb.Score, lb.Game_UUID)

    return err
}

func (s *PostgresLeaderboardStore) GetLeaderboardByGameUUID(gameUUID string) ([]*custometype.Leaderboard, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }

    defer db.Close()

    rows, err := db.Query("SELECT pseudo, score FROM leaderboard WHERE game_uuid = $1 ORDER BY score DESC;", gameUUID)

    if err != nil {
        return nil, err
    }

    var (
        pos int = 1
        leaderboards []*custometype.Leaderboard
    )

    for rows.Next() {
        p := &pos
        leaderboard := &custometype.Leaderboard{}
        rows.Scan(&leaderboard.User_UUID, &leaderboard.Score)
        leaderboard.Rank = pos
        *p = pos + 1

        leaderboards = append(leaderboards, leaderboard)
    }

    return leaderboards, nil
}

func (s *PostgresLeaderboardStore) GetTopPlayerFromGameUUID(gameUUID string, top int) ([]*custometype.Leaderboard, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }

    defer db.Close()

    rows, err := db.Query("SELECT pseudo, score FROM leaderboard WHERE game_uuid = $1 ORDER BY score DESC LIMIT $2;", gameUUID, top)

    if err != nil {
        return nil, err
    }

    var (
        pos int = 1
        leaderboards []*custometype.Leaderboard
    )

    for rows.Next() {
        p := &pos
        leaderboard := &custometype.Leaderboard{}
        rows.Scan(&leaderboard.User_UUID, &leaderboard.Score)
        leaderboard.Rank = pos
        *p = pos + 1

        leaderboards = append(leaderboards, leaderboard)
    }

    return leaderboards, nil
}
