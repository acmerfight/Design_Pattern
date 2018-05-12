package main

import (
	"fmt"
	"sync"
)

type Color struct {
	Name string
}

var (
	pool map[string]*Color
	once sync.Once
)

func init() {
	once.Do(func() {
		pool = make(map[string]*Color)
	})
}

func NewColor(name string) *Color {
	switch name {
	case "red":
		return &Color{Name: name}
	case "blue":
		return &Color{Name: name}
	case "green":
		return &Color{Name: name}
	default:
		return &Color{Name: "White"}
	}
}

func ColorFactory(name string) *Color {
	c := &Color{}
	if v, ok := pool[name]; ok {
		return v
	} else {
		fmt.Println("Create")
		c = NewColor(name)
		pool[name] = c
	}
	return c
}

func main() {
	fmt.Println(ColorFactory("red").Name)
	fmt.Println(ColorFactory("red").Name)
}
