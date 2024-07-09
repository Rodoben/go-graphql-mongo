package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodoben007/go-graphql-mongoDB/service"
)

type HRDetailsResponse struct {
	HRName string `json:"hr_name"`
	Error  string `json:"error,omitempty"`
}

var (
	gethrdetails = service.EmployeeService.GetHRDetails
)

func GetDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("body", id)
	response := HRDetailsResponse{}
	hrname, err := gethrdetails(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
	} else {
		response.HRName = hrname
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
