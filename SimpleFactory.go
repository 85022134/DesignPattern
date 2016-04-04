package main

import (
	"fmt"
)

//模式特点：工厂根据条件产生不同功能的类。
//程序实例：四则运算计算器，根据用户的输入产生相应的运算类，用这个运算类处理具体的运算。

type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() float64
}

type BaseOperation struct {
	NumA float64
	NumB float64
}

func (b *BaseOperation) SetA(a float64) {
	b.NumA = a
}

func (b *BaseOperation) SetB(bb float64) {
	b.NumB = bb
}

type AddOperation struct {
	BaseOperation
}

func (a *AddOperation) GetResult() float64 {
	return a.NumA + a.NumB
}

type SubOperation struct {
	BaseOperation
}

func (s *SubOperation) GetResult() float64 {
	return s.NumA - s.NumB
}

type OperationFactory struct {
}

func (this *OperationFactory) createOperation(operator string) (operation Operation) {
	switch operator {
	case "+":
		operation = new(AddOperation)
	case "-":
		operation = new(SubOperation)
	default:
		panic("运算符号错误！")
	}
	return
}
func SimpleFactory() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var fac OperationFactory
	operation := fac.createOperation("+")
	operation.SetA(11)
	operation.SetB(9)
	fmt.Println(operation.GetResult())
	operation = fac.createOperation("-")
	operation.SetA(11)
	operation.SetB(9)

	fmt.Println(operation.GetResult())

	operation = fac.createOperation("11")
	operation.SetA(11)
	operation.SetB(9)
}
