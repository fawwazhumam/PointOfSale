package httpclient

// import (
// 	"net/http"
// 	"pos/product-service/configs"
// 	"time"
// )

// type MerchantClient struct {
// 	urlMerchantService string
// 	httpclient *http.Client
// }

// type MerchantProductStockResponse struct {
// 	ProductID uint `json:"product_id"`
// 	TotalStock int `json:"total_stock"`
// }

// type MerchantProductStockServiceResponse struct {
// 	Message string `json:"message"`
// 	Data MerchantProductStockResponse `json:"data"`
// 	Error string `json:"error, omitempty"`
// }

// func NewMerchantClient(cfg configs.Config) MerchantClient {
// 	return &MerchantClient{
// 		httpclient: &http.Client{
// 			Timeout: 30 * time.Second,
// 		},
// 		urlMerchantService: cfg.UrlMerchantService,
// 	}
// }