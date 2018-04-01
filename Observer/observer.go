package main

import "fmt"

// 代码参考 https://blog.csdn.net/Jeanphorn/article/details/78784197 表示感谢

import (
	"fmt"
)

type Event struct {
	Data string
}

type Observer interface {
	update(event *Event)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(event *Event)
}

type ConcreteObserver struct {
	Id int64
}

func (co *ConcreteObserver) update(event *Event) {
	fmt.Println(co.Id, "update: ", event.Data)

}

type ConcreteSubject struct {
	Observers map[Observer]interface{}
}

func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.Observers[observer] = nil
}

func (cs *ConcreteSubject) Detach(observer Observer) {
	delete(cs.Observers, observer)
}

func (cs *ConcreteSubject) Notify(event *Event) {
	for observer, _ := range cs.Observers {
		observer.update(event)
	}
}

func main() {
	cs := ConcreteSubject{
		Observers: make(map[Observer]interface{}),
	}
	cs.Attach(&ConcreteObserver{1})
	cs.Attach(&ConcreteObserver{2})
	event := Event{"hello"}
	cs.Notify(&event)
}
