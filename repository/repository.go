package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/rodoben007/go-graphql-mongoDB/common"
)

type Rdms struct {
	Db *sqlx.DB
}

func Connect() *Rdms {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own

	dsn, err := loadDBproperties()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	return &Rdms{
		Db: db,
	}
}

func loadDBproperties() (string, error) {
	var dbproperties common.DBProperties
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbPropertiesBytes := []byte(os.Getenv(common.APP_DB_POSTGRES_PROPERTIES))

	err = json.Unmarshal(dbPropertiesBytes, &dbproperties)
	if err != nil {
		return "", err
	}

	dsnstr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s", dbproperties.Username, dbproperties.Dbname, dbproperties.Sslmode, dbproperties.Password, dbproperties.Host)
	return dsnstr, nil
}

func (d *Rdms) GetManagerDetails(id string) (string, error) {

	return "Managername", nil

}

func (d *Rdms) GetHRDetails(id string) (string, error) {

	return "HRrname", nil

}
