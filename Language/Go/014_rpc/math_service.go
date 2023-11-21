package main

import "errors"

type MathService struct{}

type Args struct {
	A, B int
}

type Reply struct {
	Result int
}

func (m *MathService) Add(args *Args, reply *Reply) error {
	reply.Result = args.A + args.B
	return nil
}

func (m *MathService) Divide(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	reply.Result = args.A / args.B
	return nil
}
