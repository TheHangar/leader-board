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

    _, err = db.Exec("INSERT INTO game (uuid, api_key, name) VALUES ($1, $2, $3);", g.UUID, g.ApiKey, g.Name)

    return g, err
}
func (s *PostgresGameStore) GetGames() ([]*custometype.Game, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }
    defer db.Close()

    var games []*custometype.Game

    rows, err := db.Query("SELECT api_key, uuid, name FROM game;")

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        game := custometype.Game{}
        rows.Scan(&game.ApiKey, &game.UUID, &game.Name)
        games = append(games, &game)
    }

    return games, nil
}
func (s *PostgresGameStore) GetGameByUUID(uuid string) (*custometype.Game, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }
    defer db.Close()

    var game custometype.Game

    row := db.QueryRow("SELECT uuid, api_key, name FROM game WHERE uuid=$1;", uuid)

    if err != nil {
        return nil, err
    }

    row.Scan(&game.UUID, &game.ApiKey, &game.Name)

    return &game, nil
}

func (s *PostgresGameStore) DeleteGame(uuid string) error {
    db, err := postgresConnect()

    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM game WHERE uuid = $1;", uuid)

    return err
}
