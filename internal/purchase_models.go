package internal

import "time"

type PurchaseBook struct {
	Date              *time.Time
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        int
}

type PurchaseHistory struct {
	PurchaseID        int        `json:"purchase_id"`
	Date              *time.Time `json:"date"`
	CustomerID        int        `json:"customer_id"`
	TotalPricePayment float64    `json:"total_price_payment"`
	Status            string     `json:"status"`
	PurchaseNumber    string     `json:"purchase_number"`
}

type PurchaseItem struct {
	PurchaseItemID,
	BookID,
	PurchaseHistoryID,
	Qty int
	TotalPrice float64
}

type DeletePurchaseItemsTxParams struct {
	PurchaseNumber string `json:"purchase_number"`
}

type CreatePurchaseBookTxParams struct{}
type CreatePurchaseBookTxResult struct{}

type EditBookToPurchaseParams struct {
	Qty               int
	TotalPrice        float64
	BookID            int
	PurchaseHistoryID int
	PurchaseNumber    string
}

type CreatePurchaseBookParams struct {
	// Date               *time.Time ==> auto generate from golang time.Date
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        int
}

type CreatePurchaseHistoryResult struct {
	PurchaseHistoryID int
}

type BookToPurchase struct {
	BookID            int
	Qty               int
	PurchaseHistoryID int
	CurrentStockQty   int
}

type ListBooksToPurchase []BookToPurchase

type CreateBookToPurchaseParams struct {
	BookID            int     `json:"book_id"`
	PurchaseHistoryID int     `json:"purchase_history_id"`
	Qty               int     `json:"qty"`
	TotalPrice        float64 `json:"total_price"`
	PurchaseNumber    string  `json:"purchase_number"`
}

type CreatePurchaseHistoryParams struct {
	// Date              *time.Time ==> auto generate from golang time.Date
	CustomerID        int     `json:"customer_id"`
	TotalPricePayment float64 `json:"total_price_payment"`
	Status            string  `json:"status"`          // pending or completed
	PurchaseNumber    string  `json:"purchase_number"` // PRCBOOK_20250421_RANDOMCHAR
}

type PurchaseDatatable struct {
	PurchaseNumber string     `json:"purchase_number"`
	Qty            int        `json:"qty"`
	BookTitle      string     `json:"book_title"`
	TotalPrice     float64    `json:"total_price"`
	EachPrice      float64    `json:"each_price"`
	DateOfSale     *time.Time `json:"date_of_sale"`
	Status         string     `json:"status"`
}
