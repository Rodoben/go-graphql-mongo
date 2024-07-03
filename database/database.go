package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rodoben007/go-graphql-mongoDB/common"
	"github.com/rodoben007/go-graphql-mongoDB/repository"

	"github.com/rodoben007/go-graphql-mongoDB/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Mongoclient    *mongo.Client
	PostgresClient repository.Rdms
}

var db = repository.Connect()

func loadMongoProperties() (string, error) {
	var dbproperties common.DBProperties
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbPropertiesBytes := []byte(os.Getenv(common.APP_DB_MONGO_PROPERTIES))
	err = json.Unmarshal(dbPropertiesBytes, &dbproperties)
	if err != nil {

		return "", err
	}
	dsnstr := fmt.Sprintf("%s://%s:%s@%s:%v", dbproperties.Dbname, dbproperties.Username, dbproperties.Password, dbproperties.Host, dbproperties.Port)
	return dsnstr, nil
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	uri, err := loadMongoProperties()
	if err != nil {
		log.Fatal("unable to fetch  to uri")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("unable to connect to mongo")
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil
	}
	fmt.Println("connected successfully")
	return &DB{
		Mongoclient: client,
	}

}

func (d *DB) CreateJob(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {

	jobCollection := d.Mongoclient.Database("Jobs-record").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	Insertederg, err := jobCollection.InsertOne(ctx, bson.M{
		common.JOBTITLE:       input.JobTitle,
		common.JOBDESCRIPTION: input.JobDescription,
		common.JOBCOMPANY:     input.JobCompany,
		common.JOBSPOCS: bson.M{
			common.MANAGER: bson.M{
				common.MANAGERID:   input.JobSpocs.Manager.ManagerID,
				common.MANAGERNAME: input.JobSpocs.Manager.ManagerName,
			},
			common.HUMANRESOURCE: bson.M{
				common.HRID:   input.JobSpocs.HumanResource.HrID,
				common.HRNAME: input.JobSpocs.HumanResource.HrName,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	insertedID := Insertederg.InsertedID.(primitive.ObjectID).Hex()

	return &model.JobListing{
		ID:             insertedID,
		JobTitle:       input.JobTitle,
		JobDescription: input.JobDescription,
		JobCompany:     input.JobCompany,
		JobSpocs: &model.Spocs{
			Manager: &model.Manager{
				ManagerID:   input.JobSpocs.Manager.ManagerID,
				ManagerName: input.JobSpocs.Manager.ManagerName,
			},
			HumanResource: &model.HumanResource{
				HrID:   input.JobSpocs.HumanResource.HrID,
				HrName: input.JobSpocs.HumanResource.HrName,
			},
		},
	}, nil
}

func (d *DB) UpdateJob(ctx context.Context, filter, update primitive.M, id string) (*model.JobListing, error) {

	jobCollection := d.Mongoclient.Database("Jobs-record").Collection("jobs")

	_, err := jobCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	jobListing, err := d.Job(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	return &model.JobListing{
		ID:             id,
		JobTitle:       jobListing.JobTitle,
		JobDescription: jobListing.JobDescription,
		JobCompany:     jobListing.JobCompany,
		JobSpocs: &model.Spocs{
			Manager: &model.Manager{
				ManagerID:   jobListing.JobSpocs.Manager.ManagerID,
				ManagerName: jobListing.JobSpocs.Manager.ManagerName,
			},
			HumanResource: &model.HumanResource{
				HrID:   jobListing.JobSpocs.HumanResource.HrID,
				HrName: jobListing.JobSpocs.HumanResource.HrName,
			},
		},
	}, nil

}

func (d *DB) DeleteJob(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	jobCollection := d.Mongoclient.Database("Jobs-record").Collection("jobs")

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	_, err := jobCollection.DeleteOne(ctx, filter)

	if err != nil {
		log.Println("Error deleting the id")
	}

	return &model.DeleteJobResponse{
		DeleteJobID: id,
	}, nil
}

func (d *DB) Job(ctx context.Context, id string) (*model.JobListing, error) {

	jobCollection := d.Mongoclient.Database("Jobs-record").Collection("jobs")
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing *model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return jobListing, nil
}

func (d *DB) Jobs(ctx context.Context) ([]*model.JobListing, error) {

	var joblistings []*model.JobListing
	jobCollection := d.Mongoclient.Database("Jobs-record").Collection("jobs")

	results, err := jobCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	if results.All(ctx, &joblistings); err != nil {

		log.Println("Error", err)
	}
	return joblistings, nil

}
