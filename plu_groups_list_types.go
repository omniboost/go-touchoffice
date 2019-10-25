package touchoffice

type PLUGroupsList struct {
	Status    int       `json:"status"`
	PLUGroups PLUGroups `json:"plugroups"`
}

type PLUGroups []PLUGroup

type PLUGroup struct {
	PLUgroupNumber string `json:"plugroup_number"`
	PLUgroupName   string `json:"plugroup_name"`
	AccountingCode string `json:"accounting_code"`
}
