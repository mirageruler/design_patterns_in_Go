package main

import (
	"errors"
	"time"
)

type PaymentType int

const (
	Cash PaymentType = iota
	DebitCard
	CreditCard
)

type IPaymentMethod interface {
	Init() IPaymentMethod
	Pay(amount float32) bool
}

// type IPaymentFactory interface {
// 	GetPayment(int) (IPaymentMethod, error)
// }

type CashPayment struct{}

func (c *CashPayment) Pay(amount float32) bool {
	return true
}

func (c *CashPayment) Init() IPaymentMethod {
	return c
}

type DebitCardPayment struct {
	balance float32
	expire  time.Time
}

func (d *DebitCardPayment) Pay(amount float32) bool {
	now := time.Now().In(time.UTC)
	if d.balance < amount || d.expire.Before(now) {
		return false
	}
	return true
}

func (d *DebitCardPayment) SetBalance(balance float32) IPaymentMethod {
	d.balance = balance
	return d
}

func (d *DebitCardPayment) SetExpire(date string) IPaymentMethod {
	const shortForm = "2006-01-02 15:04:05"
	d.expire, _ = time.Parse(shortForm, date)
	return d
}

func (d *DebitCardPayment) Init() IPaymentMethod {
	d.SetBalance(100)
	d.SetExpire("2030-12-30 23:59:59")
	return d
}

type CreditCardPayment struct {
	balance   float32
	tolerance float32
	expire    time.Time
}

func (cc *CreditCardPayment) Pay(amount float32) bool {
	now := time.Now().In(time.UTC)
	if cc.balance+cc.tolerance < amount || cc.expire.Before(now) {
		return false
	}
	return true
}

func (cc *CreditCardPayment) Init() IPaymentMethod {
	cc.SetBalance(100)
	cc.SetTolerance(200)
	cc.SetExpire("2030-12-30 23:59:59")
	return cc
}

func (cc *CreditCardPayment) SetTolerance(tolerance float32) IPaymentMethod {
	cc.tolerance = tolerance
	return cc
}

func (cc *CreditCardPayment) SetBalance(balance float32) IPaymentMethod {
	cc.balance = balance
	return cc
}

func (cc *CreditCardPayment) SetExpire(date string) IPaymentMethod {
	const shortForm = "2006-01-02 15:04:05"
	cc.expire, _ = time.Parse(shortForm, date)
	return cc
}

type PaymentFactory struct{}

func (pF *PaymentFactory) GetPayment(t PaymentType) (IPaymentMethod, error) {
	switch t {
	case Cash:
		return new(CashPayment).Init(), nil
	case DebitCard:
		return new(DebitCardPayment).Init(), nil
	case CreditCard:
		return new(CreditCardPayment).Init(), nil
	default:
		return nil, errors.New("invalid payment type")
	}
}
