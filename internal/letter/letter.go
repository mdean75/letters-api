package letter

import "time"

type Letter struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	TO        string    `json:"to" bson:"to"`
	From      string    `json:"from" bson:"from"`
	Message   string    `json:"message" bson:"msg"`
	CreatedTs time.Time `json:"createdTs" bson:"createdTs"`
	User      string    `json:"user" bson:"user"`
}
