package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID    primitive.ObjectID `json:"id"`
	Title string             `json:"title"`
	Body  string             `json:"body"`
}
