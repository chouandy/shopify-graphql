package shopifygraphql

// MailingAddressInput mailing address input struct
type MailingAddressInput struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Company   string `json:"company,omitempty"`
	Country   string `json:"country,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Province  string `json:"province,omitempty"`
	Zip       string `json:"zip,omitempty"`
}
