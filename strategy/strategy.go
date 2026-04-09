package strategy

import "fmt"

type PaymentStrategyBaseClass interface {
	Pay(amount float64)
}

type UpiPayment struct{}

func (u *UpiPayment) Pay(amount float64) {
	fmt.Printf("Sending payment of %f via UPI", amount)
}

type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) {
	fmt.Printf("Sending payment of %f via Card", amount)
}

type ShoppingCart struct {
	PaymentStrategyBaseClass
}

func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategyBaseClass) {
	s.PaymentStrategyBaseClass = paymentStrategy
}

func (s *ShoppingCart) Checkout(amount float64) {
	if s.PaymentStrategyBaseClass == nil {
		fmt.Println("Illegal stuff, base class not initialized")
		return
	}
	s.PaymentStrategyBaseClass.Pay(amount)
}
