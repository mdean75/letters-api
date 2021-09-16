package config

import (
	"os"
	"strconv"
)

type EnvVar struct {
	dbConn    string
	oktaToken string
}

func (e EnvVar) LoadConfig() Configuration {
	e.dbConn = os.Getenv("dbconn")
	e.oktaToken = os.Getenv("oktaToken")

	dc := os.Getenv("DEBUG_CORS")
	debugCors, err := strconv.ParseBool(dc)
	if err != nil {
		return NewConfiguration(NewMongoConfig(e.dbConn), NewOktaConfig(e.oktaToken), false)
	}

	return NewConfiguration(NewMongoConfig(e.dbConn), NewOktaConfig(e.oktaToken), debugCors)
}
