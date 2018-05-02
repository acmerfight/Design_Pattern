package main

import "fmt"

type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct{}

type KoreanGirlFriendFactory struct{}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func getGirlFriend(typeGf string) GirlFriend {
	factoryMap := map[string]AbstractFactory{
		"Indian": IndianGirlFriendFactory{},
		"Korean": KoreanGirlFriendFactory{},
	}
	return factoryMap[typeGf].CreateMyLove()
}

func main() {
	a := getGirlFriend("Indian")
	fmt.Println(a.eyesColor)
	b := getGirlFriend("Korean")
	fmt.Println(b.eyesColor)

}
