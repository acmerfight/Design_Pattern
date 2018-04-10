package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// 代码参考 https://coolshell.cn/articles/17929.html

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			start := time.Now()
			out = targetFunc.Call(in)
			fmt.Println("elapsedTime: ", time.Since(start))
			return
		})

	decoratedFunc.Set(v)
	return
}

func test() {
	fmt.Println(math.Pow(2, 100))
}

type MyTest func()

func main() {
	var myTest MyTest
	Decorator(&myTest, test)
	myTest()
}
