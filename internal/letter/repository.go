package letter

import (
	"fmt"
	"os"

	"letters-api/internal/config"
	"letters-api/internal/db"
)

type Repository interface {
	Insert(l Letter) (string, error)
	FetchOne(id string) (Letter, error)
	FetchAllForUser(user string) ([]Letter, error)
}

type Controller struct {
	Datastore Repository
	DebugCors bool
}

func NewController(r Repository, debugCors bool) *Controller {
	return &Controller{Datastore: r, DebugCors: debugCors}
}

func CreateController() *Controller {
	c := config.EnvVar{}.LoadConfig()
	if c.DBConn() == "" {
		fmt.Println("dbconn is not set")
		os.Exit(100)
	}

	mongoConn, err := db.NewMongoDatabase(c.DBConn())
	if err != nil {
		fmt.Println(err)
	}

	dao := NewDAO(mongoConn, "l3o", "savedLetter")

	return NewController(dao, c.DebugCors)
}
