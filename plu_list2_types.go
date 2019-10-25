package touchoffice

type PLUList2 struct {
	Products Products `json:"products"`
	Status   int      `json:"status"`
}

type Products []Product

type Product struct {
	Plu            string  `json:"plu"`
	Name           string  `json:"name"`
	Randomcode     string  `json:"randomcode"`
	Plugroup       int     `json:"plugroup"`
	Department     int     `json:"department"`
	Taxrate        int     `json:"taxrate"`
	Price1L1       float64 `json:"price1l1"`
	Price2L1       float64 `json:"price2l1"`
	Price3L1       int     `json:"price3l1"`
	Plugroupname   string  `json:"plugroupname"`
	Departmentname string  `json:"departmentname"`
}
