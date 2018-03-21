package main

import "fmt"

type Processor interface {
	step1()
	step2()
	step3()
}

type Abstract struct {
	processor Processor
}

func (Abstract) step2() {
	fmt.Println("step2")
}

func (Abstract) process(pro Processor) {
	pro.step1()
	pro.step2()
	pro.step3()
}

type Concrete struct {
	abstract Abstract
}

func (Concrete) step1() {
	fmt.Println("step1")
}

func (Concrete) step3() {
	fmt.Println("step3")
}

func main() {
	abstract := Abstract{}
	concrete := Concrete{abstract}
	concrete.abstract.process()

