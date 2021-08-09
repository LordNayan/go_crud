package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(task *Task) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	task.ID = primitive.NewObjectID()

	//Create a handle to the respective collection in the database.
	collection := client.Database("tasks").Collection("tasks")
	//Perform InsertOne operation & validate against the error.
	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Printf("Could not create Task: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func GetByTitle(title string) (Task, error) {
	result := Task{}
	filter := bson.D{primitive.E{Key: "title", Value: title}}
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	//Create a handle to the respective collection in the database.
	collection := client.Database("tasks").Collection("tasks")

	//Perform FindOne operation & validate against the error.
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func DeleteByID(ID primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "id", Value: ID}}
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	//Create a handle to the respective collection in the database.
	collection := client.Database("tasks").Collection("tasks")

	//Perform DeleteOne operation & validate against the error.
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return result without any error.
	return nil
}

func UpdateByID(task *Task) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "id", Value: task.ID}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: task.Title},
		primitive.E{Key: "body", Value: task.Body},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	//Create a handle to the respective collection in the database.
	collection := client.Database("tasks").Collection("tasks")

	//Perform UpdateOne operation & validate against the error.
	_, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
