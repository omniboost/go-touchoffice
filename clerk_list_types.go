package touchoffice

type ClerkList struct {
	Status int              `json:"status"`
	Clerks map[string]Clerk `json:"clerks"`
}

type Clerk struct {
	ClerkNumber   string `json:"clerk_number"`
	ClerkName     string `json:"clerk_name"`
	IbuttonNumber string `json:"ibutton_number"`
	SecretNumber  string `json:"secret_number"`
}
