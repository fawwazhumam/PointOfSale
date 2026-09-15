package response

import (
	"pos/warehouse-service/pkg/pagination"
	"time"
)

type WarehouseResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Photo     string    `json:"photo"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllWarehousesResponse struct {
	Warehouses []WarehouseResponse           `json:"warehouses"`
	Pagination pagination.PaginationResponse `json:"pagination"`
}
