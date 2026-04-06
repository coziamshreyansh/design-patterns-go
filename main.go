package main

import (
	"design-patterns-go/abstractFactory"
	"design-patterns-go/adapter"
	"design-patterns-go/builder"
	"design-patterns-go/constants"
	"design-patterns-go/decorator"
	"design-patterns-go/factory"
	"design-patterns-go/singleton"
	"fmt"
)

func factoryPatternDriver() {
	//object using factory
	smsObj, err := factory.Create(constants.Sms)
	if err != nil {
		fmt.Print("Something went wrong!")
	}

	smsObj.Send("99074", "Heyyyyaaaaaa!")

	emailObj, err := factory.Create(constants.Email)
	if err != nil {
		fmt.Print("Something went wrong!")
	}
	emailObj.Send("shreyyy@gmail.com", "Heyyy!")
}

func abstractFactoryPatternDriver() {
	var factory abstractFactory.UIFactorInterface
	darkTheme := true
	if darkTheme {
		factory = abstractFactory.DarkTheme{}
	} else {
		factory = abstractFactory.LightTheme{}
	}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckBox()

	button.Render()
	checkbox.Render()
}

func singletonPatternDriver() {
	instance := singleton.GetInstance()
	instance.Info("Success")
	instance.Error("Failure")
}

func builderPatternDriver() {
	builderInstance := builder.CreateUserBuilderInstance()
	builderInstance.WithMobile("678900000")
	builderInstance.WithEmail("shre@gmi.l")
	builderInstance.WithName("shreyansh")
	user := builderInstance.Build()
	fmt.Println("User object:- ", user)
}

func adapterPatternDriver() {
	existingService := adapter.ExistingPaymentService{}
	var paymentAdapter adapter.PaymentAdapter
	paymentAdapter = adapter.NewPaymentServiceAdapter(existingService)
	paymentAdapter.ProcessPayment(100000)
}

func decoratorPatternDriver() {
	base := &decorator.MessageNotifier{}
	smsDecorator := &decorator.SmsDecorator{NotifierDecorator: decorator.NotifierDecorator{Wrapped: base}}
	emailDecorator := &decorator.EmailDecorator{NotifierDecorator: decorator.NotifierDecorator{Wrapped: smsDecorator}}
	emailDecorator.Send("Heyyyy!")
}

func main() {
	// basically all creational patterns are a way of writing if else in diff manner
	// factoryPatternDriver()
	// abstractFactoryPatternDriver()
	// singletonPatternDriver()
	// builderPatternDriver()
	// adapterPatternDriver()
	decoratorPatternDriver()
}
