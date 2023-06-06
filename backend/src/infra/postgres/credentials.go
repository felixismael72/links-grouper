package postgres

import (
	"errors"
	"fmt"
)

var (
	host     string
	port     int
	user     string
	password string
	dbName   string
)

func SetupCredentials(newHost string, newPort int, newUser, newPassword, newDBName string) {
	host = newHost
	port = newPort
	user = newUser
	password = newPassword
	dbName = newDBName
}

func getFullConnectionURI() (string, error) {
	if !hasValidCredentials() {
		return "", errors.New("invalid credentials for the postgres database")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName), nil
}

func hasValidCredentials() bool {
	return user != "" && password != "" && host != "" && port > 0
}
