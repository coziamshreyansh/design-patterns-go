package main

import (
	"design-patterns-go/constants"
	"design-patterns-go/factory"
	"fmt"
)

func main() {
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
