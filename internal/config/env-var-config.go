package config

import (
	"os"
	"strconv"
)

type EnvVar struct {
	dbConn            string
	oktaToken         string
	loginRadiusApiKey string
}

func (e EnvVar) LoadConfig() Configuration {
	e.dbConn = os.Getenv("dbconn")
	e.oktaToken = os.Getenv("oktaToken")
	e.loginRadiusApiKey = os.Getenv("LOGIN_RADIUS_API_KEY")

	dc := os.Getenv("DEBUG_CORS")
	debugCors, err := strconv.ParseBool(dc)
	if err != nil {
		return NewConfiguration(NewMongoConfig(e.dbConn), NewOktaConfig(e.oktaToken), NewLoginRadiusConfig(e.loginRadiusApiKey), false)
	}

	return NewConfiguration(NewMongoConfig(e.dbConn), NewOktaConfig(e.oktaToken), NewLoginRadiusConfig(e.loginRadiusApiKey), debugCors)
}
