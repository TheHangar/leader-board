package store

import (
	custometype "github.com/TheHangar/leader-board/custome_type"
	_ "github.com/lib/pq"
)

type PostgresAdminUserStore struct {
}

func NewPostgresUserStore() *PostgresAdminUserStore {
    return &PostgresAdminUserStore{}
}

func (s *PostgresAdminUserStore) AddAdminUser(user *custometype.User) (*custometype.User, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }

    defer db.Close()

    return nil, nil
}

func (s *PostgresAdminUserStore) GetAdminUserByUsername(username string) (*custometype.User, error) {
    db, err := postgresConnect()

    if err != nil {
        return nil, err
    }

    row := db.QueryRow("SELECT * FROM admin WHERE username = $1;", username)

    if err != nil {
        return nil, err
    }

    var user custometype.User

    err = row.Scan(&user.ID, &user.Username, &user.Password)

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *PostgresAdminUserStore) DeleteAdminUserByUsername(username string) error {
    return nil
}
