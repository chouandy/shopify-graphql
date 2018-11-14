package shopifygraphql

import "github.com/shopspring/decimal"

// ProductVariantInput product variant input struct
type ProductVariantInput struct {
	Barcode              string                `json:"barcode,omitempty"`
	CompareAtPrice       *decimal.Decimal      `json:"compareAtPrice,omitempty"`
	FulfillmentServiceID string                `json:"fulfillmentServiceId,omitempty"`
	HarmonizedSystemCode string                `json:"harmonizedSystemCode,omitempty"`
	ID                   string                `json:"id,omitempty"`
	ImageID              string                `json:"imageId,omitempty"`
	ImageSrc             string                `json:"imageSrc,omitempty"`
	InventoryManagement  string                `json:"inventoryManagement,omitempty"`
	InventoryPolicy      string                `json:"inventoryPolicy,omitempty"`
	InventoryQuantities  []InventoryLevelInput `json:"inventoryQuantities,omitempty"`
	Metafields           []MetafieldInput      `json:"metafields,omitempty"`
	Options              []string              `json:"options,omitempty"`
	Position             int                   `json:"position,omitempty"`
	Price                *decimal.Decimal      `json:"price,omitempty"`
	ProductID            string                `json:"productId,omitempty"`
	RequireShipping      *bool                 `json:"requiresShipping,omitempty"`
	Sku                  string                `json:"sku,omitempty"`
	TaxCode              string                `json:"taxCode,omitempty"`
	Taxable              *bool                 `json:"taxable,omitempty"`
	Title                string                `json:"title,omitempty"`
	Weight               *decimal.Decimal      `json:"weight,omitempty"`
	WeightUnit           string                `json:"weightUnit,omitempty"`
}
