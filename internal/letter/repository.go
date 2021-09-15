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
}

type Controller struct {
	Datastore Repository
}

func NewController(r Repository) *Controller {
	return &Controller{Datastore: r}
}

func CreateController() *Controller {
	c := config.EnvVar{}.LoadConfig()

	mongoConn, err := db.NewMongoDatabase(c.DBConn())
	if err != nil {
		fmt.Println(err)
	}

	dao := NewDAO(mongoConn, "l3o", "savedLetter")

	return NewController(dao)
}
