package touchoffice

type ReceiptsList struct {
	Receipts Receipts `json:"Receipts"`
	Status   int      `json:"status"`
}

type Receipts []Receipt

type Receipt struct {
	Body   []string `json:"body"`
	Header struct {
		Mode              int      `json:"mode"`
		Saleid            string   `json:"saleid"`
		Site              int      `json:"site"`
		Till              int      `json:"till"`
		Tillname          string   `json:"tillname"`
		Datetime          string   `json:"datetime"`
		Ordernumber       int      `json:"ordernumber"`
		Consecutivenumber int      `json:"consecutivenumber"`
		Clerknumber       int      `json:"clerknumber"`
		Clerkname         string   `json:"clerkname"`
		Saletotal         float64  `json:"saletotal"`
		Itemcount         int      `json:"itemcount"`
		Check             int      `json:"check"`
		Table             int      `json:"table"`
		Covers            int      `json:"covers"`
		Location          string   `json:"location"`
		Finalisekeys      []string `json:"finalisekeys"`
		Transactionkeys   int      `json:"transactionkeys"`
		Customer          Customer `json:"customer"`
	} `json:"header"`
	Items ReceiptItems `json:"items"`
}

type Customer struct {
	Customernumber int         `json:"customernumber"`
	Customername   interface{} `json:"customername"`
}

type ReceiptItems []ReceiptItem

type ReceiptItem struct {
	Plu   int     `json:"plu"`
	Item  string  `json:"item"`
	Qty   int     `json:"qty"`
	Value float64 `json:"value"`
}
