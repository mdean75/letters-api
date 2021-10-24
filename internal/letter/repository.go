package letter

import (
	"fmt"

	"letters-api/internal/config"
	"letters-api/internal/db"
)

type Repository interface {
	Insert(l Letter) (string, error)
	FetchOne(id string) (Letter, error)
	FetchAllForUser(user string) ([]Letter, error)
	FetchMetaForUser(user string) ([]Letter, error)
	FetchMetaForLetter(id string) (Letter, error)
	FetchContent(id string) (interface{}, error)

	InsertHTML(h Data) (string, error)
}

type Controller struct {
	Datastore         Repository
	LoginRadiusApiKey string
	DebugCors         bool
	Done              chan struct{}
}

func NewController(r Repository, apiKey string, debugCors bool, ch chan struct{}) *Controller {
	return &Controller{Datastore: r, LoginRadiusApiKey: apiKey, DebugCors: debugCors, Done: ch}
}

func CreateController(ch chan struct{}) *Controller {
	c := config.EnvVar{}.LoadConfig()
	if c.DBConn() == "" {
		fmt.Println("dbconn is not set")
		ch <- struct{}{}
		return nil
	}

	mongoConn, err := db.NewMongoDatabase(c.DBConn())
	if err != nil {
		fmt.Println(err)
		ch <- struct{}{}
		return nil
	}

	dao := NewDAO(mongoConn, "l3o", "savedLetter")

	return NewController(dao, c.LoginRadiusConfig.ApiKey(), c.DebugCors, ch)
}
