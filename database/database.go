package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB representa a conexão com o banco de dados
type DB struct {
	*sql.DB
}

// InitDB inicializa a conexão com o banco de dados
func InitDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Criar a tabela drivers se não existir
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS drivers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        acceleration REAL,
        braking REAL,
        sharp_turn BOOLEAN
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
