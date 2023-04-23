package main

import "fmt"

// Complex subsystem parts
type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
