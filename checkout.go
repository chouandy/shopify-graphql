package shopifygraphql

import "github.com/shurcooL/graphql"

// CheckoutCreateInput checkout create input
type CheckoutCreateInput struct {
	Email                 graphql.String           `json:"email,omitempty"`
	LineItems             []*CheckoutLineItemInput `json:"lineItems,omitempty"`
	CustomAttributes      []*AttributeInput        `json:"customAttributes,omitempty"`
	AllowPartialAddresses graphql.Boolean          `json:"allowPartialAddresses,omitempty"`
	Note                  graphql.String           `json:"note,omitempty"`
	ShippingAddress       *MailingAddressInput     `json:"shippingAddress,omitempty"`
}

// CheckoutLineItemInput checkout line item input
type CheckoutLineItemInput struct {
	VariantID        graphql.ID        `json:"variantId,omitempty"`
	Quantity         graphql.Int       `json:"quantity,omitempty"`
	CustomAttributes []*AttributeInput `json:"customAttributes,omitempty"`
}
