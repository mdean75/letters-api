package letter

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c *MongoDataStore) InsertHTML(h Data) (string, error) {
	opts := options.BucketOptions{
		//Name:           &name,
	}
	bucket, err := gridfs.NewBucket(c.db, &opts)
	if err != nil {
		return "", err
	}

	r := strings.NewReader(h.Content)

	uploadopts := options.UploadOptions{Metadata: map[string]interface{}{
		"user":      h.User,
		"to":        h.To,
		"from":      h.From,
		"createdTs": h.CreatedTs,
		"title":     h.Title,
	}}
	stream, err := bucket.UploadFromStream(fmt.Sprintf("%s-%s.html", h.User, h.CreatedTs), r, &uploadopts)
	if err != nil {
		return "", err
	}

	log.Printf("Write file to DB was successful. \n")

	return stream.Hex(), nil
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

func (c *MongoDataStore) FetchContent(id string) (interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Letter{}, err
	}

	//c.db.Collection("fs.files").FindOne(context.TODO(), bson.M{"_id": _id})
	bucket, err := gridfs.NewBucket(c.db)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = bucket.DownloadToStream(_id, &buf)
	if err != nil {
		return nil, err
	}

	return buf.String(), nil
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

func (c *MongoDataStore) FetchMetaForUser(user string) ([]Letter, error) {
	cur, err := c.db.Collection("fs.files").Find(context.TODO(), bson.M{"metadata.user": user})
	if err != nil {
		return nil, err
	}

	var letters []DAOLetterMeta
	err = cur.All(context.TODO(), &letters)
	if err != nil {
		return nil, err
	}

	var retLetters []Letter
	for _, l := range letters {
		retLetters = append(retLetters, NewLetter(l.ID.Hex(), l.Metadata.TO, l.Metadata.From, l.Metadata.User, l.Metadata.Title, l.Metadata.CreatedTs))
	}

	return retLetters, nil
}

func (c *MongoDataStore) FetchMetaForLetter(id string) (Letter, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Letter{}, err
	}

	sr := c.db.Collection("fs.files").FindOne(context.TODO(), bson.M{"_id": _id})
	if sr.Err() != nil {
		return Letter{}, err
	}

	var l DAOLetterMeta
	err = sr.Decode(&l)
	if err != nil {
		return Letter{}, err
	}

	//var retLetters []Letter
	//for _, l := range letters {
	//	retLetters = append(retLetters, NewLetter(l.ID.Hex(), l.Metadata.TO, l.Metadata.From, l.Metadata.User, l.Metadata.Title, l.Metadata.CreatedTs))
	//}

	return NewLetter(l.ID.Hex(), l.Metadata.TO, l.Metadata.From, l.Metadata.User, l.Metadata.Title, l.Metadata.CreatedTs), nil
}
