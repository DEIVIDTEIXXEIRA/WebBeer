package data

import (
	"database/sql"
	"webbeer/config"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringDeConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
