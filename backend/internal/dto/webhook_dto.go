package dto

import "github.com/google/uuid"

type WebhookPayload struct {
	PurchasingId uuid.UUID     `json:"purchasing_id"`
	UserID       uuid.UUID     `json:"user_id"`
	SupplierID   uuid.UUID     `json:"supplier_id"`
	GrandTotal   int64         `json:"grand_total"`
	Items        []ItemPayload `json:"items"`
}

type ItemPayload struct {
	ItemID    uuid.UUID `json:"item_id"`
	ItemName  string    `json:"item_name"`
	ItemPrice int64     `json:"item_price"`
	Qty       int64     `json:"qty"`
	SubTotal  int64     `json:"sub_total"`
}
