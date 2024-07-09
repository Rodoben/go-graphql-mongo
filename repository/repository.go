package repository

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rodoben007/go-graphql-mongoDB/common"
	"gorm.io/gorm"
)

type Rdms struct {
	Db *gorm.DB
}
type SalesOrderStatus struct {
	ID         int    `json:"id"`
	StatusID   string `json:"status_id"`
	StatusName string `json:"status_name"`
}

var SalesOrderRepository EmployeeDetails = &Rdms{Db: Connect().Db}

func (d *Rdms) GetManagerDetails(id string) (string, error) {

	return "Managername", nil

}

func (d *Rdms) GetHRDetails(ctx context.Context, id string) (string, error) {

	var result SalesOrderStatus
	fmt.Println("I am here in repo", d)
	if err := d.Db.Raw(common.Query, id).Scan(&result).Error; err != nil {
		return "", err
	}
	fmt.Println("____", result)
	return result.StatusName, nil

}
