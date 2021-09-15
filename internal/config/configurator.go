package config

type DBConfigurator interface {
	DBConn() string
	SetDBConn(s string)
}

type OktaConfigurator interface {
	APIToken() string
	SetAPIToken(s string)
}
