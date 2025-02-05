package database

import (
	"context"
	"fmt"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertIntoDB(jobPostings []common.JobPosting) error {
	client := GetDBClient()
	if client == nil {
		return fmt.Errorf("error: unable to get db client connection")
	}

	database := client.Database("jobsDB")
	collection := database.Collection("jobsCollection")

	return InsertNewJobPostings(collection, jobPostings)
}

func InsertNewJobPostings(collection *mongo.Collection, jobPostings []common.JobPosting) error {
	for _, job := range jobPostings {
		doc := bson.M{
			"company":       job.Company,
			"jobId":         job.JobId,
			"title":         job.JobTitle,
			"locationsText": job.Location,
			"postedOn":      job.PostedOn,
			"externalPath":  job.ExternalPath,
			"insertedOn":    time.Now(),
		}

		_, err := collection.InsertOne(context.TODO(), doc)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				fmt.Println(err.Error())
				continue
			} else {
				return fmt.Errorf("error while inserting data into db: %v", err.Error())
			}
		}
	}

	return nil
}

func DeleteJobsFromDB(jobPostings []common.JobPosting) error {
	client := GetDBClient()
	if client == nil {
		return fmt.Errorf("error: unable to get db client connection")
	}

	database := client.Database("jobsDB")
	collection := database.Collection("jobsCollection")

	var err error
	for _, job := range jobPostings {
		er := deleteJobById(collection, job.JobId)
		if er != nil {
			err = er
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func deleteJobById(collection *mongo.Collection, jobId string) error {
	filter := bson.M{"jobId": jobId}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error while deleting job from db: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no job found with jobId: %s", jobId)
	}

	fmt.Printf("Successfully deleted job with jobId: %s\n", jobId)
	return nil
}
