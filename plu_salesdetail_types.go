package touchoffice

type SalesList struct {
	Sales  Sales `json:"Sales"`
	Status int   `json:"status"`
}

type Sales []Sale

type Sale struct {
	Header struct {
		Saleid            string   `json:"saleid"`
		Site              int      `json:"site"`
		Till              int      `json:"till"`
		Date              string   `json:"date"`
		Time              string   `json:"time"`
		Ordernumber       int      `json:"ordernumber"`
		Consecutivenumber int      `json:"consecutivenumber"`
		Clerknumber       int      `json:"clerknumber"`
		Saletotal         float64  `json:"saletotal"`
		Itemcount         int      `json:"itemcount"`
		Check             int      `json:"check"`
		Table             int      `json:"table"`
		Covers            int      `json:"covers"`
		Location          string   `json:"location"`
		Finalisekeys      []string `json:"finalisekeys"`
		Transactionkeys   []string `json:"transactionkeys"`
	} `json:"header"`
	SalesItems SalesItems `json:"items"`
}

type SalesItems []SalesItem

type SalesItem struct {
	Plu        int     `json:"plu"`
	Item       string  `json:"item"`
	Department string  `json:"department"`
	PluGroup   string  `json:"plu_group"`
	Qty        int     `json:"qty"`
	Value      float64 `json:"value"`
}
