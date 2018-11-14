package shopifygraphql

// CheckoutCreateInput checkout create input struct
type CheckoutCreateInput struct {
	Email                 string                  `json:"email,omitempty"`
	LineItems             []CheckoutLineItemInput `json:"lineItems,omitempty"`
	CustomAttributes      []AttributeInput        `json:"customAttributes,omitempty"`
	AllowPartialAddresses bool                    `json:"allowPartialAddresses,omitempty"`
	Note                  string                  `json:"note,omitempty"`
	ShippingAddress       *MailingAddressInput    `json:"shippingAddress,omitempty"`
}

// CheckoutLineItemInput checkout line item input struct
type CheckoutLineItemInput struct {
	VariantID        string           `json:"variantId,omitempty"`
	Quantity         int              `json:"quantity,omitempty"`
	CustomAttributes []AttributeInput `json:"customAttributes,omitempty"`
}
