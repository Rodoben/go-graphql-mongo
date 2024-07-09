package service

import "context"

type EmployeeDetailsService interface {
	GetHRDetails(ctx context.Context, id string) (string, error)
}
