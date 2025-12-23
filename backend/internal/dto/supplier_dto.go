package dto

type CreateSupplierRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
type UpdateSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
}

type SupplierResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
