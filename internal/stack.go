package internal

import "fmt"

type Stack struct {
	values []float64
}

func NewStack() *Stack {
	var stack = &Stack{}
	stack.values = make([]float64, 0, 10)
	return stack
}

func (this *Stack) Len() int {
	return len(this.values)
}

type ValueHandler func(value float64)

func (this *Stack) Foreach(handler ValueHandler) {
	for _, v := range this.values {
		handler(v)
	}
}

func (this *Stack) new() {
	this.values = make([]float64, 0, 10)
}

func (this *Stack) Push(v float64) {
	this.values = append(this.values, v)
}

func (this *Stack) Pop() float64 {
	var len = this.Len()
	var result = this.values[len - 1]
	this.values = this.values[:len - 1]
	return result
}

func (this *Stack) Drop() bool {
	var len = this.Len()
	if len == 0 {
		return false
	}

	this.values = this.values[:len - 1]
	return true
}

func (this *Stack) Print() {
	fmt.Println(this.values)
}

func (this *Stack) Negate() bool {
	if this.Len() == 0 {
		return false
	}

	var value = this.Pop()
	this.Push(-value)
	return true
}

func (this *Stack) Duplicate() bool {
	if this.Len() == 0 {
		return false
	}

	var value = this.Pop()
	this.Push(value)
	this.Push(value)
	return true
}

func (this *Stack) Swap() bool {
	if this.Len() < 2 {
		return false
	}

	var x = this.Pop()
	var y = this.Pop()
	this.Push(x)
	this.Push(y)
	return true
}

func (this *Stack) Add() bool {
	if this.Len() < 2 {
		return false
	}

	var y = this.Pop()
	var x = this.Pop()
	this.Push(x + y)
	return true
}

func (this *Stack) Substract() bool {
	if this.Len() < 2 {
		return false
	}

	var y = this.Pop()
	var x = this.Pop()
	this.Push(x - y)
	return true
}

func (this *Stack) Multiply() bool {
	if this.Len() < 2 {
		return false
	}

	var y = this.Pop()
	var x = this.Pop()
	this.Push(x * y)
	return true
}

func (this *Stack) Divide() bool {
	if this.Len() < 2 {
		return false
	}

	var y = this.Pop()
	var x = this.Pop()
	this.Push(x / y)
	return true
}
