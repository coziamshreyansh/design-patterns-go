package main

import (
	"design-patterns-go/abstractFactory"
	"design-patterns-go/constants"
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
func main() {
	// basically all creational patterns are a way of writing if else in diff manner
	// factoryPatternDriver()
	// abstractFactoryPatternDriver()
	singletonPatternDriver()
}
