package common

const (
	JOBTITLE       = "jobtitle"
	JOBDESCRIPTION = "jobdescription"
	JOBCOMPANY     = "jobCompany"
	JOBSPOCS       = "jobSpocs"
	MANAGER        = "manager"
	MANAGERID      = "managerId"
	MANAGERNAME    = "managerName"
	HUMANRESOURCE  = "humanResource"
	HRID           = "hrID"
	HRNAME         = "hrName"
)

const (
	APP_DB_POSTGRES_PROPERTIES = "APP_DB_POSTGRES_PROPERTIES"
	APP_DB_MONGO_PROPERTIES    = "APP_DB_MONGO_PROPERTIES"
)

type DBProperties struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Dbname       string `json:"dbname"`
	Sslmode      string `json:"sslmode"`
	ReadOnlyHost string `json:"readOnlyHost"`
	ReadOnlyPort int    `json:"readOnlyPort"`
}

const (
	Query = "select id, status_id, status_name from sales_order.sales_order_status where status_id =$1;"
)
