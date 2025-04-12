package adapterdp

// import "fmt"

// // PaymentProcessor defines a common interface for payment processors
// type PaymentProcessor interface {
// 	Pay(amount float64) string
// }

// // PayPalProcessor is the built-in payment implementation
// type PayPalProcessor struct{}

// func (p *PayPalProcessor) Pay(amount float64) string {
// 	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
// }

// // StripeX is a third-party SDK we can't modify
// type StripeX struct{}

// func (s *StripeX) MakePayment(cents int) string {
// 	return fmt.Sprintf("Processed payment of %d cents with StripeX", cents)
// }

// // StripeXAdapter adapts StripeX to conform to PaymentProcessor interface
// type StripeXAdapter struct {
// 	stripe *StripeX
// }

// // NewStripeXAdapter returns a new StripeXAdapter
// func NewStripeXAdapter(stripe *StripeX) *StripeXAdapter {
// 	return &StripeXAdapter{stripe: stripe}
// }

// func (a *StripeXAdapter) Pay(amount float64) string {
// 	cents := int(amount * 100)
// 	return a.stripe.MakePayment(cents)
// }
