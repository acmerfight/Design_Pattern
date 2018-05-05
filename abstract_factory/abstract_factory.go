package main

import "fmt"

// this is my concrete girl friend
type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type Baby struct {
	nationality string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
	CreateBaby() Baby
}

type IndianFactory struct{}

type KoreanFactory struct{}

func (a IndianFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a IndianFactory) CreateBaby() Baby {
	return Baby{"Hindi"}
}

func (a KoreanFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func (a KoreanFactory) CreateBaby() Baby {
	return Baby{"Korean"}
}

func getFacotry(typeGf string) interface{} {
	factoryMap := map[string]AbstractFactory{
		"Indian": IndianFactory{},
		"Korean": KoreanFactory{},
	}
	return factoryMap[typeGf]
}

func main() {
	var factory interface{}
	factory = getFacotry("Indian")
	fmt.Println(factory.(IndianFactory).CreateMyLove().eyesColor)
	fmt.Println(factory.(IndianFactory).CreateBaby().nationality)
	factory = getFacotry("Korean")
	fmt.Println(factory.(KoreanFactory).CreateMyLove().eyesColor)
	fmt.Println(factory.(KoreanFactory).CreateBaby().nationality)
}
