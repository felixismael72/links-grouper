package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type ConnectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ ConnectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	connectionURI, err := getFullConnectionURI()
	if err != nil {
		return nil, err
	}

	return sqlx.Connect("postgres", connectionURI)
}

func (DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Fatal().Err(err)
	}
}
