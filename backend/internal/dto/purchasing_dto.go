package dto

type CreatePurchasingRequest struct {
	UserID     string                    `json:"user_id"`
	Date       string                    `json:"date" validate:"required, date"`
	SupplierID string                    `json:"supplier_id" validate:"required"`
	Items      []PurchasingDetailRequest `json:"items"`
}

type PurchasingDetailRequest struct {
	ItemID string `json:"item_id"`
	Qty    int64  `json:"qty"`
}
