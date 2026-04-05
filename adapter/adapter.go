package adapter

import "fmt"

type ExistingPaymentService struct{}

// existing payment service can only process in paise
func (e *ExistingPaymentService) MakePaymentInPaise(amountInPaise float64) {
	fmt.Printf("Payment of %f processed in paise", amountInPaise)
}

// but client will give me amt in rupees
// so i will write adapter for conv rs to paise for new clients
// so that the existing functionality will not break, if any old clients are using

type PaymentAdapter interface {
	ProcessPayment(amount float64)
}

type ExistingPaymentServiceAdapter struct {
	existingService ExistingPaymentService
}

func NewPaymentServiceAdapter(existingService ExistingPaymentService) PaymentAdapter {
	return ExistingPaymentServiceAdapter{existingService: existingService}
}

func (e ExistingPaymentServiceAdapter) ProcessPayment(amountInRupees float64) {
	// conv rs to paise
	amount := amountInRupees * 100
	e.existingService.MakePaymentInPaise(amount)
	fmt.Printf("Processed payment in Rupees using adapter")
}
