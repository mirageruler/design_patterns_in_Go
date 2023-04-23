package main

import (
	"fmt"
	"log"
	"strings"
)

type OsType int

const (
	WINDOWS OsType = iota
	MAC
)

type IGUIFactory interface {
	createButton() IButton
	createCheckbox() ICheckbox
}

type IButton interface {
	paint() string
}

type ICheckbox interface {
	paint() string
}

type WinButton struct{}

func (b *WinButton) paint() string {
	return "From WinButton"
}

type WinCheckbox struct{}

func (c *WinCheckbox) paint() string {
	return "From WinCheckbox"
}

type WinFactory struct{}

func (wf *WinFactory) createButton() IButton {
	return new(WinButton)
}

func (wf *WinFactory) createCheckbox() ICheckbox {
	return new(WinCheckbox)
}

type MacButton struct{}

func (b *MacButton) paint() string {
	return "From MacButton"
}

type MacCheckbox struct{}

func (c *MacCheckbox) paint() string {
	return "From MacCheckbox"
}

type MacFactory struct{}

func (mf *MacFactory) createButton() IButton {
	return new(MacButton)
}

func (mf *MacFactory) createCheckbox() ICheckbox {
	return new(MacCheckbox)
}

type Application struct {
	factory IGUIFactory
}

func (a *Application) SetOsFactory(os OsType) (*Application, error) {
	switch os {
	case WINDOWS:
		a.factory = new(WinFactory)
	case MAC:
		a.factory = new(MacFactory)
	default:
		return nil, fmt.Errorf("wrong OS type passed")
	}

	return a, nil
}

func (a *Application) SetFactory(f IGUIFactory) {
	a.factory = f
}

func (a *Application) GetFactory() IGUIFactory {
	return a.factory
}

func (a *Application) CreateUI() string {
	button := a.factory.createButton()
	checkBox := a.factory.createCheckbox()
	log.Println("Rendering UI...")
	b, c := button.paint(), checkBox.paint()
	return strings.Join([]string{b, c}, "-")
}
