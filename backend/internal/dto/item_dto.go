package dto

type CreateItemRequest struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

type ItemResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

type UpdateItemRequest struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}
