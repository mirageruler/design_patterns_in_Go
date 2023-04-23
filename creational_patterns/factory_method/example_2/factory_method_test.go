package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {
}

func (s *TestSuite) TestCashPay() {
	cashPayment := new(CashPayment)
	assert.Equal(s.T(), true, cashPayment.Pay(100))
}

func (s *TestSuite) TestDebitCardPay() {
	debitCardPayment := new(DebitCardPayment).Init()
	assert.Equal(s.T(), true, debitCardPayment.Pay(100))
}

func (s *TestSuite) TestDebitCardInsufficientBalance() {
	debitCardPayment := new(DebitCardPayment)
	debitCardPayment.SetBalance(50)
	debitCardPayment.SetExpire("2030-12-30 23:59:59")
	assert.Equal(s.T(), false, debitCardPayment.Pay(100))
}

func (s *TestSuite) TestDebitCardPayExpire() {
	debitCardPayment := new(DebitCardPayment)
	debitCardPayment.SetBalance(100)
	now := time.Now().UTC()
	nowSubOneMinute := now.Add(time.Duration(-1) * time.Minute)
	nowAddOneMinute := now.Add(time.Duration(1) * time.Minute)
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	endOfToday := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	type parameters struct {
		expire time.Time
	}
	testCases := []struct {
		name   string
		args   parameters
		expect bool
	}{
		{
			name:   "Yesterday",
			args:   parameters{expire: yesterday},
			expect: false,
		},
		{
			name:   "Now",
			args:   parameters{expire: now},
			expect: false,
		},
		{
			name:   "One minute before Now",
			args:   parameters{expire: nowSubOneMinute},
			expect: false,
		},
		{
			name:   "One minute after Now",
			args:   parameters{expire: nowAddOneMinute},
			expect: true,
		},
		{
			name:   "End of Today",
			args:   parameters{expire: endOfToday},
			expect: true,
		},
		{
			name:   "Tomorrow",
			args:   parameters{expire: tomorrow},
			expect: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			debitCardPayment.SetExpire(tc.args.expire.Format("2006-01-02 15:04:05"))
			assert.Equal(s.T(), tc.expect, debitCardPayment.Pay(100))
		})
	}
}

func (s *TestSuite) TestGetPayment() {
	paymentFactory := new(PaymentFactory)
	debitPayment, _ := paymentFactory.GetPayment(DebitCard)
	assert.Equal(s.T(), reflect.TypeOf(&DebitCardPayment{}), reflect.TypeOf(debitPayment))

	cashPayment, _ := paymentFactory.GetPayment(Cash)
	assert.Equal(s.T(), reflect.TypeOf(&CashPayment{}), reflect.TypeOf(cashPayment))

	_, err := paymentFactory.GetPayment(PaymentType(99))
	assert.Equal(s.T(), err.Error(), "invalid payment type")
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
