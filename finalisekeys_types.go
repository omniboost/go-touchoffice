package touchoffice

type FinalisekeysList struct {
	Finalisekeys Finalisekeys `json:"sales"`
	Status       int          `json:"status"`
}

type Finalisekeys []Finalisekey

type Finalisekey struct {
	Record   int     `json:"record"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Value    float64 `json:"value"`
	Flag     int     `json:"flag"`
}
