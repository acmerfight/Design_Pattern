package main

import "fmt"

//这个不是真正的 template method ，因为调用 process 执行的是 Abstract 的 step2，step2 不能访问 Concrete 的数据。
//如果要像 Python 一样实现真正的 template method 可以参考 http://hackthology.com/object-oriented-inheritance-in-go.html

// Golang 的 sort 也是 template method 的一种好的实现 https://golang.org/src/sort/sort.go

type Processor interface {
	step1()
	step2()
	step3()
}

type Abstract struct {
	pro Processor
}

func (abs *Abstract) step1() {
	if abs.pro == nil {
		return
	} else {
		abs.pro.step1()
	}
}

func (abs Abstract) step3() {
	if abs.pro == nil {
		return
	} else {
		abs.pro.step3()
	}
}

func (abs *Abstract) step2() {
	fmt.Println("step2")
}

func (abs *Abstract) process() {
	abs.step1()
	abs.step2()
	abs.step3()
}

type Concrete struct {
	Abstract
}

func (con *Concrete) step1() {
	fmt.Println("step1")
}

func (con *Concrete) step3() {
	fmt.Println("step3")
}

func main() {
	abstract := &Abstract{}
	abstract.pro = &Concrete{}
	abstract.process()
}
