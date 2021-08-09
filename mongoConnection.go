package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := goDotEnvVariable("MONGODB_USERNAME")
	password := goDotEnvVariable("MONGODB_PASSWORD")
	clusterEndpoint := goDotEnvVariable("MONGODB_ENDPOINT")
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	if err == nil {
		fmt.Println("Connected to MongoDB!")
	}
	return client, ctx, cancel
}
