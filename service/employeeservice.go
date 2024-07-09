package service

import (
	"context"

	"github.com/rodoben007/go-graphql-mongoDB/repository"
)

type EmployeeDetailsImpl struct{}

var (
	EmployeeService EmployeeDetailsService = EmployeeDetailsImpl{}
	gethrdetails                           = repository.SalesOrderRepository.GetHRDetails
)

func (e EmployeeDetailsImpl) GetHRDetails(ctx context.Context, id string) (string, error) {
	return gethrdetails(ctx, id)
}
