package graph

import (
	"context"
	"fmt"

	"github.com/rodoben007/go-graphql-mongoDB/database"
	"github.com/rodoben007/go-graphql-mongoDB/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var db = database.Connect()

type Resolver struct {
	Db database.DB
}

func (r *Resolver) CreateJob(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	r.Db = *db
	managerName, err := r.Db.GetManagerDetails(input.JobSpocs.Manager.ManagerID)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch getmangerDetails()")
	}

	hrName, err := r.Db.GetHRDetails(input.JobSpocs.HumanResource.HrID)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch getHRDetails()")
	}
	input.JobSpocs.Manager.ManagerName = managerName
	input.JobSpocs.HumanResource.HrName = hrName

	resp, err := r.Db.CreateJob(ctx, input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (r *Resolver) UpdateJob(ctx context.Context, id string, input *model.UpdateJobListingInput) (*model.JobListing, error) {
	r.Db = *db
	updateJobInfo := bson.M{}
	if input.JobTitle != "" {
		updateJobInfo["title"] = input.JobTitle
	}
	if input.JobCompany != "" {
		updateJobInfo["jobcompany"] = input.JobCompany
	}
	if input.JobDescription != "" {
		updateJobInfo["jobdescription"] = input.JobDescription
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}
	resp, err := r.Db.UpdateJob(ctx, filter, update, id)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (r *Resolver) DeleteJob(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	r.Db = *db

	resp, err := r.Db.DeleteJob(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (r *Resolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	r.Db = *db
	resp, err := r.Db.Job(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch Job")
	}
	return resp, nil
}

func (r *Resolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	r.Db = *db
	resp, err := r.Db.Jobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch Jobs")
	}
	return resp, nil
}
