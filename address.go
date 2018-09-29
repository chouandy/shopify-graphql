package shopifygraphql

import "github.com/shurcooL/graphql"

// MailingAddressInput mailing address input
type MailingAddressInput struct {
	Address1  graphql.String `json:"address1,omitempty"`
	Address2  graphql.String `json:"address2,omitempty"`
	City      graphql.String `json:"city,omitempty"`
	Company   graphql.String `json:"company,omitempty"`
	Country   graphql.String `json:"country,omitempty"`
	FirstName graphql.String `json:"firstName,omitempty"`
	LastName  graphql.String `json:"lastName,omitempty"`
	Phone     graphql.String `json:"phone,omitempty"`
	Province  graphql.String `json:"province,omitempty"`
	Zip       graphql.String `json:"zip,omitempty"`
}
