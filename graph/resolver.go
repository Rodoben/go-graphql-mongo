package graph

import (
	"context"
	"fmt"

	"github.com/rodoben007/go-graphql-mongoDB/database"
	"github.com/rodoben007/go-graphql-mongoDB/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var db = database.Connect()

type Resolver struct {
	Db database.DB
}

func (r *Resolver) CreateJob(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {

	fmt.Println("I am createjob")

	r.Db = *db
	return &model.JobListing{
		ID:       "1234",
		JobTitle: "ronaldddd",
	}, nil
}
