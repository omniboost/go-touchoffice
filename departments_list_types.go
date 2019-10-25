package touchoffice

type DepartmentsList struct {
	Status      int         `json:"status"`
	Departments Departments `json:"departments"`
}

type Departments []Department

type Department struct {
	DepartmentNumber string `json:"department_number"`
	DepartmentName   string `json:"department_name"`
	AccountingCode   string `json:"accounting_code"`
}
