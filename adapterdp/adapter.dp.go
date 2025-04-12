package adapterdp

import (
	"fmt"
)

// Paypal is the concrete implementation
type Paypal struct{}

func (p *Paypal) MakePayment(amount float64, currency string) {
	fmt.Println("Paypal payment proccess")
}

// Stripe is the concrete implementation
type Stripe struct{}

func (s *Stripe) ChargeAmount(amount float64, currency string) {
	fmt.Println("Stripe payment proccess")
}

// target interface expected by the client
type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string)
	// RollbackProcess(amount float64)
}

// ShoppingChart is client that use PaymentProcessor()
type ShoppingChart struct {
	processor PaymentProcessor
}

func (c *ShoppingChart) Checkout(amount float64, currency string) {
	c.processor.ProcessPayment(amount, currency)
}

func NewShoppingChart(processor PaymentProcessor) ShoppingChart {
	return ShoppingChart{
		processor,
	}
}

// this adapter implement PaymentProcessor interface
type PaymentAdapter struct {
	PaymentMethod any
}

func (a *PaymentAdapter) ProcessPayment(amount float64, currency string) {
	switch method := a.PaymentMethod.(type) {
	case *Paypal:
		// implementation details tetep di method tiap payment method
		method.MakePayment(amount, currency)
	case *Stripe:
		method.ChargeAmount(amount, currency)
	default:
		fmt.Println("Unsupported payment method")
	}
}

func NewPaymentAdapter(argPaymentMethod any) *PaymentAdapter {
	return &PaymentAdapter{
		PaymentMethod: argPaymentMethod,
	}
}
