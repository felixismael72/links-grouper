package config

type API struct {
	Server Server
	PG     PG
}

type Server struct {
	Host string
	Port int
}

type PG struct {
	Username string
	Password string
	DBName   string
	Server   Server
}
