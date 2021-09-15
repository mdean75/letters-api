package config

import "os"

type EnvVar struct {
	dbConn    string
	oktaToken string
}

func (e EnvVar) LoadConfig() Configuration {
	e.dbConn = os.Getenv("dbconn")
	e.oktaToken = os.Getenv("oktaToken")

	return NewConfiguration(NewMongoConfig(e.dbConn), NewOktaConfig(e.oktaToken))
}
