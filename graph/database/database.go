package database

import (
	"context"
	"log"

	"github.com/lunatictiol/graphql-go/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"time"
)

var connectionString = ""

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

func (db *DB) CreateJobListing(input model.CreateJobListingInput) *model.JobListing {
	return nil
}
func (db *DB) UpdateJobListing(id string, input model.UpdateJobListingInput) *model.JobListing {
	return nil
}
func (db *DB) DeleteJobListing(id string) *model.DeleteJobResponse {
	return nil
}
func (db *DB) GetAllJobs() []*model.JobListing {
	return nil
}
func (db *DB) GetJobByID(id string) *model.JobListing {
	return nil
}
