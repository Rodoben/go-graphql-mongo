package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/rodoben007/go-graphql-mongoDB/graph/model"
)

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	//panic(fmt.Errorf("not implemented: CreateJob - createJob"))

	return r.Resolver.CreateJob(ctx, input)

}

// UpdateJob is the resolver for the updateJob field.
func (r *mutationResolver) UpdateJob(ctx context.Context, id string, input *model.UpdateJobListingInput) (*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: UpdateJob - updateJob"))
}

// DeleteJob is the resolver for the deleteJob field.
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteJob - deleteJob"))
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: Jobs - jobs"))
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: Job - job"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct {
	*Resolver
}
type queryResolver struct{ *Resolver }
