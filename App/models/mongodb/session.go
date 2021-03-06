package mongodb

import (
	"context"
	"go-grpc/App/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username"`
	IsActive  bool          `json:"is_active" bson:"is_active"`
	ExpiredAt time.Time     `json:"expired_at" bson:"expired_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

func CreateSession(username string, expiredAt time.Time) error {
	var sessionCollection = db.OpenCollection(db.Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var session Session

	session.ExpiredAt = expiredAt
	session.Username = username
	session.IsActive = true
	session.CreatedAt = time.Now()

	_, err := sessionCollection.InsertOne(ctx, session)

	return err
}

func SessionActive(username string) (bool, error) {
	var sessionCollection = db.OpenCollection(db.Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var session Session

	filter := bson.M{
		"is_active": bson.M{"$eq": true},
		"username":  bson.M{"$eq": username},
	}

	err := sessionCollection.FindOne(ctx, filter).Decode(&session)

	// Token expired
	if time.Now().After(session.ExpiredAt) {
		return false, nil
	}

	return session.IsActive, err
}

func UpdateSession(username string) error {
	var sessionCollection = db.OpenCollection(db.Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"username": bson.M{"$eq": username}}
	update := bson.M{"$set": bson.M{"is_active": false}}

	_, err := sessionCollection.UpdateMany(ctx, filter, update)

	return err
}
