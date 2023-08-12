package store

import (
	"database/sql"
	"fmt"
	"os"

	custometype "github.com/TheHangar/leader-board/custome_type"
    _ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type LeaderboardStore interface {
    AddLeaderboard(*custometype.Leaderboard) (*custometype.Leaderboard, error)
    GetLeaderboardByGameUUID(string) ([]*custometype.Leaderboard, error)
    GetTopPlayerFromGameUUID(uuid string, top int) ([]*custometype.Leaderboard, error)
}

type GameStore interface {
    AddGame(*custometype.Game) (*custometype.Game, error)
    GetGames() ([]*custometype.Game, error)
    DeleteGame(uuid string) error
}

type AdminUserStore interface {
    AddAdminUser(*custometype.User) (*custometype.User, error)
    GetAdminUserByUsername(string) (*custometype.User, error)
    DeleteAdminUserByUsername(string) error
}

type Storage struct {
    User AdminUserStore
    Game GameStore
    Leaderboard LeaderboardStore
}

func NewPostgresStore() *Storage {
    us := NewPostgresUserStore()
    gs := NewPostgresGameStore()
    ls := NewPostgresLeaderboardStore()
    return &Storage{ User: us, Game: gs, Leaderboard: ls }
}

func postgresConnect() (*sql.DB, error) {
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"))
    db, err := sql.Open("postgres", connectionString)
	if err != nil {
        return nil, err
	}

    return db, nil
}

func PostgresTest() error {
    db, err := postgresConnect()

    if err != nil {
        return err
    }
    defer db.Close()

    return db.Ping()
}

func PostgresSeed() error {
    db, err := postgresConnect()

    if err != nil {
        return err
    }
    defer db.Close()

    createTables := `
        CREATE TABLE IF NOT EXISTS game (
            id SERIAL,
            uuid VARCHAR(256) PRIMARY KEY,
            name VARCHAR(256)
        );
        CREATE TABLE IF NOT EXISTS admin (
            id SERIAL PRIMARY KEY,
            username VARCHAR(256),
            password VARCHAR(256)
        );
        CREATE TABLE IF NOT EXISTS leaderboard (
            id SERIAL PRIMARY KEY,
            game_uuid VARCHAR(256) REFERENCES game(uuid),
            pseudo VARCHAR(256),
            score FLOAT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
    
    _, err = db.Exec(createTables)

    if err != nil {
        return err
    }
    hashPwd, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PWD")), bcrypt.DefaultCost)

    if err != nil {
        return err
    }

    _, err = db.Exec("INSERT INTO admin(username, password) VALUES ($1, $2)", os.Getenv("ADMIN_NAME"), string(hashPwd))

    return err
}
