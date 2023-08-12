package store

import custometype "github.com/TheHangar/leader-board/custome_type"

type PostgresLeaderboardStore struct {
}

func NewPostgresLeaderboardStore() *PostgresLeaderboardStore {
    return &PostgresLeaderboardStore{}
}

func (s *PostgresLeaderboardStore) AddLeaderboard(*custometype.Leaderboard) (*custometype.Leaderboard, error) {
    return nil, nil
}

func (s *PostgresLeaderboardStore) GetLeaderboardByGameUUID(gameUUID string) ([]*custometype.Leaderboard, error) {
    return nil, nil
}

func (s *PostgresLeaderboardStore) GetTopPlayerFromGameUUID(gameUUID string, top int) ([]*custometype.Leaderboard, error) {
    return nil, nil
}
