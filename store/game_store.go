package store

import (
	custometype "github.com/TheHangar/leader-board/custome_type"
)

type PostgresGameStore struct {
}

func NewPostgresGameStore() *PostgresGameStore {
    return &PostgresGameStore{}
}

func (s *PostgresGameStore) AddGame(g *custometype.Game) (*custometype.Game, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO game (uuid, name) VALUES ($1, $2);", g.UUID, g.Name)

    return g, err
}
func (s *PostgresGameStore) GetGames() ([]*custometype.Game, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }
    defer db.Close()

    var games []*custometype.Game

    rows, err := db.Query("SELECT uuid, name FROM game;")

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        game := custometype.Game{}
        rows.Scan(&game.UUID, &game.Name)
        games = append(games, &game)
    }

    return games, nil
}
func  (s *PostgresGameStore)DeleteGame(uuid string) error {
    return nil
}
