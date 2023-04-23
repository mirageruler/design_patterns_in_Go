package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFactory struct {
	mock.Mock
}

func (m *MockFactory) createButton() IButton {
	args := m.Called()
	return args.Get(0).(IButton)
}

func (m *MockFactory) createCheckbox() ICheckbox {
	args := m.Called()
	return args.Get(0).(ICheckbox)
}

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {}

func (s *TestSuite) TestWinFactoryButton() {
	winFactory := new(WinFactory)
	button := winFactory.createButton()
	assert.Equal(s.T(), "*main.WinButton", reflect.TypeOf(button).String())
}

func (s *TestSuite) TestWinFactoryCreateCheckbox() {
	winFactory := new(WinFactory)
	checkBox := winFactory.createCheckbox()
	assert.Equal(s.T(), "*main.WinCheckbox", reflect.TypeOf(checkBox).String())
}

func (s *TestSuite) TestWinButtonPaint() {
	button := new(WinButton)
	str := button.paint()
	assert.Equal(s.T(), "From WinButton", str)
}

func (s *TestSuite) TestWinCheckboxPaint() {
	checkBox := new(WinCheckbox)
	str := checkBox.paint()
	assert.Equal(s.T(), "From WinCheckbox", str)
}

func (s *TestSuite) TestMacFactoryButton() {
	macFactory := new(MacFactory)
	button := macFactory.createButton()
	assert.Equal(s.T(), "*main.MacButton", reflect.TypeOf(button).String())
}

func (s *TestSuite) TestMacFactoryCreateCheckbox() {
	macFactory := new(MacFactory)
	checkBox := macFactory.createCheckbox()
	assert.Equal(s.T(), "*main.MacCheckbox", reflect.TypeOf(checkBox).String())
}

func (s *TestSuite) TestMacButtonPaint() {
	button := new(MacButton)
	str := button.paint()
	assert.Equal(s.T(), "From MacButton", str)
}

func (s *TestSuite) TestMacCheckboxPaint() {
	checkBox := new(MacCheckbox)
	str := checkBox.paint()
	assert.Equal(s.T(), "From MacCheckbox", str)
}

func (s *TestSuite) TestApplicationSetOsFactory() {
	winFactory, _ := new(Application).SetOsFactory(WINDOWS)
	assert.Equal(s.T(), "*main.WinFactory", reflect.TypeOf(winFactory.GetFactory()).String())
	macFactory, _ := new(Application).SetOsFactory(MAC)
	assert.Equal(s.T(), "*main.MacFactory", reflect.TypeOf(macFactory.GetFactory()).String())
	_, err := new(Application).SetOsFactory(OsType(99))
	assert.Equal(s.T(), err.Error(), "wrong OS type passed")
}

func (s *TestSuite) TestApplicationCreateUI() {
	mockFactory := new(MockFactory)
	app := new(Application)
	app.SetFactory(mockFactory)
	mockFactory.On("createButton").Return(&WinButton{})
	mockFactory.On("createCheckbox").Return(&WinCheckbox{})
	render := app.CreateUI()
	assert.Equal(s.T(), render, "From WinButton-From WinCheckbox")
	mockFactory.AssertNumberOfCalls(s.T(), "createButton", 1)
	mockFactory.AssertNumberOfCalls(s.T(), "createCheckbox", 1)
	mockFactory.AssertExpectations(s.T())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
