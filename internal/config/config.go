// Package config defines the configuration details for the application. Use of interfaces provides for different types
// of configuration such as env variables or json file.
package config

// Loader is the interface implemented by types that provide a means to load config values.
type Loader interface {
	LoadConfig() Configuration
}

// Configuration holds the configuration model for the application
type Configuration struct {
	MongoConfig
	OktaConfig
	DebugCors bool
}

func NewConfiguration(m MongoConfig, o OktaConfig, debugCors bool) Configuration {
	return Configuration{m, o, debugCors}
}

type MongoConfig struct {
	dbConn string
}

func NewMongoConfig(conn string) MongoConfig {
	return MongoConfig{dbConn: conn}
}

func (m *MongoConfig) DBConn() string {
	return m.dbConn
}

func (m *MongoConfig) SetDBConn(dbConn string) {
	m.dbConn = dbConn
}

type OktaConfig struct {
	apiToken string
}

func NewOktaConfig(token string) OktaConfig {
	return OktaConfig{apiToken: token}
}

func (o *OktaConfig) APIToken() string {
	return o.apiToken
}

func (o *OktaConfig) SetAPIToken(token string) {
	o.apiToken = token
}
