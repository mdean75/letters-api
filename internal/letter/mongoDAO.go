package letter

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"letters-api/internal/db"
)

type MongoDataStore struct {
	*db.MongoConn
	db  *mongo.Database
	col *mongo.Collection
}

func NewDAO(dbconn *db.MongoConn, db, col string) Repository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

func (c *MongoDataStore) Insert(l Letter) (string, error) {
	res, err := c.col.InsertOne(context.TODO(), l)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", res.InsertedID), nil
}

func (c *MongoDataStore) FetchOne(id string) (Letter, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Letter{}, err
	}
	res := c.col.FindOne(context.TODO(), bson.M{"_id": _id})
	if res.Err() != nil {
		return Letter{}, res.Err()
	}

	var letter Letter

	err = res.Decode(&letter)
	if err != nil {
		return Letter{}, err
	}

	return letter, nil
}

func (c *MongoDataStore) FetchAllForUser(user string) ([]Letter, error) {
	cur, err := c.col.Find(context.TODO(), bson.M{"user": user})
	if err != nil {
		return nil, err
	}

	var letters []Letter
	err = cur.All(context.TODO(), &letters)
	if err != nil {
		return nil, err
	}

	return letters, nil
}
