package letter

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Letter struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	TO   string `json:"to" bson:"to"`
	From string `json:"from" bson:"from"`
	//Message   string    `json:"message" bson:"msg"`
	CreatedTs time.Time `json:"createdTs" bson:"createdTs"`
	User      string    `json:"user" bson:"user"`
	Title     string    `json:"title" bson:"title"`
	//Metadata Metadata `json:"metadata" bson:"metadata"`
}

func NewLetter(id, to, from, user, title string, createdTs time.Time) Letter {
	return Letter{
		ID:        id,
		TO:        to,
		From:      from,
		CreatedTs: createdTs,
		User:      user,
		Title:     title,
	}
}

type DAOLetterMeta struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TO        string             `json:"to" bson:"to"`
	From      string             `json:"from" bson:"from"`
	Message   string             `json:"message" bson:"msg"`
	CreatedTs time.Time          `json:"createdTs" bson:"createdTs"`
	User      string             `json:"user" bson:"user"`
	Title     string             `json:"title" bson:"title"`
	Metadata  Metadata           `json:"metadata" bson:"metadata"`
}

type Metadata struct {
	TO        string    `json:"to" bson:"to"`
	From      string    `json:"from" bson:"from"`
	Message   string    `json:"message" bson:"msg"`
	CreatedTs time.Time `json:"createdTs" bson:"createdTs"`
	User      string    `json:"user" bson:"user"`
	Title     string    `json:"title" bson:"title"`
}
