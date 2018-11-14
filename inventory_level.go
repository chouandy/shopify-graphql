package shopifygraphql

// InventoryLevelInput inventory level input struct
type InventoryLevelInput struct {
	AvailableQuantity int    `json:"availableQuantity,omitempty"`
	LocationID        string `json:"locationId,omitempty"`
}
