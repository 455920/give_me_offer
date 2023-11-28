package main

import "errors"

type MathService struct{}

type Request struct {
	A, B int
}

type Respond struct {
	Result int
}

func (m *MathService) Add(rqst *Request, rsp *Respond) error {
	rsp.Result = rqst.A + rqst.B
	return nil
}

func (m *MathService) Divide(rqst *Request, rsp *Respond) error {
	if rqst.B == 0 {
		return errors.New("division by zero")
	}
	rsp.Result = rqst.A / rqst.B
	return nil
}
