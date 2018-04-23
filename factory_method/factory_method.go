package main

import (
	"errors"
	"fmt"
)

const (
	STOVE = iota
	FRIDGE
)

type Appliance interface {
	Start()
	GetPurpose()
}

func CreateAppliance(t int) (Appliance, error) {
	switch t {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	default:
		return nil, errors.New("invalid Appliance Type")
	}
}

type Stove struct {
	typeName string
}

func (sv *Stove) Start() {
	sv.typeName = " Stove "
}

func (sv *Stove) GetPurpose() {
	fmt.Println("I am a " + sv.typeName + " I cook food!!")
}

type Fridge struct {
	typeName string
}

func (fr *Fridge) Start() {
	fr.typeName = " Fridge "
}

func (fr *Fridge) GetPurpose() {
	fmt.Println("I am a " + fr.typeName + " I cool stuff down!!")
}

func main() {
	appliance, _ := CreateAppliance(1)
	appliance.Start()
	appliance.GetPurpose()
	appliance, _ = CreateAppliance(0)
	appliance.Start()
	appliance.GetPurpose()
}
