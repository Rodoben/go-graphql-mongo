package routes

import (
	"net/http"

	han "github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/rodoben007/go-graphql-mongoDB/database"
	"github.com/rodoben007/go-graphql-mongoDB/graph"
	"github.com/rodoben007/go-graphql-mongoDB/handler"
	"github.com/rodoben007/go-graphql-mongoDB/repository"
)

var (
	getDetails = handler.GetDetails
)

func RestRoutes(router *mux.Router) {

	router.HandleFunc("/gethrdetails/{id}", getDetails).Methods(http.MethodGet)

}

func GraphRoutes(router *mux.Router) {
	srv := han.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Db: database.DB{
			Mongoclient:    database.Connect().Mongoclient,
			PostgresClient: *repository.Connect(),
		},
	}}))
	router.Handle("/graph", srv)
}
