package repository

import "context"

type EmployeeDetails interface {
	GetHRDetails(ctx context.Context, id string) (string, error)
}
