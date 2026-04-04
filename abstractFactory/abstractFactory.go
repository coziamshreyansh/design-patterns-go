package abstractFactory

import "fmt"

// abstractFactory -> its just a factory pattern with collection of objects not just one

// dark theme -> dark button and check box objects wil be created

// light theme -> dark button and check box objects wil be created

type Button interface {
	Render()
}
type CheckBox interface {
	Render()
}

type LightButton struct{}
type LightCheckBox struct{}

func (l LightButton) Render() {
	fmt.Printf("Rendering light button")
}
func (l LightCheckBox) Render() {
	fmt.Printf("Rendering light checkbox")
}

type DarkButton struct{}
type DarkCheckBox struct{}

func (d DarkButton) Render() {
	fmt.Printf("Rendering dark button")
}
func (d DarkCheckBox) Render() {
	fmt.Printf("Rendering dark checkbox")
}

type UIFactorInterface interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

type LightTheme struct{}

func (l LightTheme) CreateButton() Button {
	return LightButton{}
}

func (l LightTheme) CreateCheckBox() CheckBox {
	return LightCheckBox{}
}

type DarkTheme struct{}

func (d DarkTheme) CreateButton() Button {
	return DarkButton{}
}

func (d DarkTheme) CreateCheckBox() CheckBox {
	return DarkCheckBox{}
}
