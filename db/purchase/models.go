package db

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
