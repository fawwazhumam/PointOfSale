package request

type CreateWarehouseProductRequest struct {
	WarehouseID uint `json:"warehouse_id" validate:"required"`
	ProductID   uint `json:"product_id" validate:"required"`
	Stock       int  `json:"stock" validate:"min=0"`
}

type GetAllWarehouseProductRequest struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Search    string `query:"search"`
	SortBy    string `query:"sort_by"`
	SortOrder string `query:"sort_order"`
}
