package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alert struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId"`
	Crypto      string             `bson:"crypto"`
	TargetPrice float64            `bson:"targetPrice"`
	Email       string             `bson:"email"`
	CreatedAt   int64              `bson:"createdAt"`
}