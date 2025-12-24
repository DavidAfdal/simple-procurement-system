package dto

type CreatePurchasingRequest struct {
	UserID     string                    `json:"user_id"`
	Date       string                    `json:"date" validate:"required,date"`
	SupplierID string                    `json:"supplier_id" validate:"required"`
	Items      []PurchasingDetailRequest `json:"items"`
}

type PurchasingDetailRequest struct {
	ItemID string `json:"item_id"`
	Qty    int64  `json:"qty"`
}

type PurchasingResponse struct {
	ID         string                     `json:"id"`
	Date       string                     `json:"date"`
	UserID     string                     `json:"user_id"`
	SupplierID string                     `json:"supplier_id"`
	GrandTotal int64                      `json:"grand_total"`
	Items      []PurchasingDetailResponse `json:"items"`
}

type PurchasingDetailResponse struct {
	ItemID    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice int64  `json:"item_price"`
	Qty       int64  `json:"qty"`
	SubTotal  int64  `json:"sub_total"`
}
