package request

type CreateWarehouseRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Photo   string `json:"photo"`
	Phone   string `json:"phone" validate:"required"`
}

type GetAllWarehouseRequest struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Search    string `query:"search"`
	SortBy    string `query:"sort_by"`
	SortOrder string `query:"sort_order"`
}
