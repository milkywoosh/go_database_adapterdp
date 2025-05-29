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
	PurchaseID        int
	Date              *time.Time
	CustomerID        int
	TotalPricePayment float64
	Status            string
	PurchaseNumber    string
}

type PurchaseItem struct {
	PurchaseItemID,
	BookID,
	PurchaseHistoryID,
	Qty int
	TotalPrice float64
}

type DeletePurchaseItemsTxParams struct {
	PurchaseNumber string
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
}

type BookToPurchase struct {
	BookID            int
	Qty               int
	PurchaseHistoryID int
	CurrentStockQty   int
}

type ListBooksToPurchase []BookToPurchase

type CreateBookToPurchaseParams struct {
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        float64
	PurchaseNumber    string
}

type CreatePurchaseHistoryParams struct {
	// Date              *time.Time ==> auto generate from golang time.Date
	CustomerID        int
	TotalPricePayment float64
	Status            string // pending or completed
	PurchaseNumber    string // PRCBOOK_20250421_RANDOMCHAR
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
