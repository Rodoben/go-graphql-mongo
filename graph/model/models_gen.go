// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListingInput struct {
	JobTitle       string      `json:"jobTitle"`
	JobDescription string      `json:"jobDescription"`
	JobCompany     string      `json:"jobCompany"`
	JobSpocs       *SpocsInput `json:"jobSpocs"`
}

type DeleteJobResponse struct {
	DeleteJobID string `json:"deleteJobId"`
}

type HumanResource struct {
	HrID   string `json:"hrID"`
	HrName string `json:"hrName"`
}

type HumanResourceInput struct {
	HrID   string `json:"hrID"`
	HrName string `json:"hrName"`
}

type JobListing struct {
	ID             string `json:"id"`
	JobTitle       string `json:"jobTitle"`
	JobDescription string `json:"jobDescription"`
	JobCompany     string `json:"jobCompany"`
	JobSpocs       *Spocs `json:"jobSpocs"`
}

type Manager struct {
	ManagerID   string `json:"managerID"`
	ManagerName string `json:"managerName"`
}

type ManagerInput struct {
	ManagerID   string `json:"managerID"`
	ManagerName string `json:"managerName"`
}

type Mutation struct {
}

type Query struct {
}

type Spocs struct {
	Manager       *Manager       `json:"manager"`
	HumanResource *HumanResource `json:"humanResource"`
}

type SpocsInput struct {
	Manager       *ManagerInput       `json:"manager"`
	HumanResource *HumanResourceInput `json:"humanResource"`
}

type UpdateJobListingInput struct {
	JobTitle       string      `json:"jobTitle"`
	JobDescription string      `json:"jobDescription"`
	JobCompany     string      `json:"jobCompany"`
	JobSpocs       *SpocsInput `json:"jobSpocs"`
}
