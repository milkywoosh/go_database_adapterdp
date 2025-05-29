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
