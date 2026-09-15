package response

import "pos/warehouse-service/pkg/pagination"

type WarehouseProductResponse struct {
	ID            uint   `json:"id"`
	WarehouseID   uint   `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name,omitempty"`
	ProductID     uint   `json:"product_id"`
	Stock         int    `json:"stock"`
}

type GetAllWarehouseProductsResponse struct {
	WarehouseProducts []WarehouseProductResponse    `json:"warehouse_products"`
	Pagination        pagination.PaginationResponse `json:"pagination"`
}
