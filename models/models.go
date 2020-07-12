package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
}