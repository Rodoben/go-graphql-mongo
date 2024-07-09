mutation createJob($input: CreateJobListingInput!) {
  createJob(input: $input) {
    id
    jobTitle
    jobDescription
    jobCompany
  }
}

{
  "input": {
    "jobTitle": "ronald",
    "jobDescription": "ffvcfdc",
    "jobCompany": "harman",
    "jobSpocs": {
      "manager": {
        "managerID": 123,
        "managerName": "ron"
      },
      "humanResource": {
        "hrID": 1234,
        "hrName": "123"
      }
    }
  }
}



mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) 
{ updateJob(id:$id,input:$input)
  {
    jobCompany
  } 
}

{ 
  "id": "667e71622aa5925161da7f00", 
   "input": {
    "jobTitle": "asfdfronaldfvfdsv",
    "jobDescription": "ffvcfdc",
    "jobCompany": "harman",
    "jobSpocs": {
      "manager": {
        "managerID": 123,
        "managerName": "ron"
      },
      "humanResource": {
        "hrID": 1234,
        "hrName": "123"
      }
    }
  }
}