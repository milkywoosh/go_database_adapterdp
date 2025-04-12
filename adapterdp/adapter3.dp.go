package adapterdp

import "fmt"

type IPaymentProcessor interface {
	Pay(amount float64) string
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

type StripeX struct{}

// SDK type should be implement using adapter
func (s *StripeX) MakePayment(cents int) string {
	return fmt.Sprintf("Processed payment of %d cents with StripeX", cents)
}

type StripeXAdapter struct {
	Adaptee *StripeX
}

func NewStripeXAdapter(Adaptee *StripeX) *StripeXAdapter {
	return &StripeXAdapter{
		Adaptee: Adaptee,
	}
}

func (p StripeXAdapter) Pay(amount float64) string {
	return p.Adaptee.MakePayment(int(amount))
}
