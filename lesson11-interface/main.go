package main

import (
	"fmt"

	"quoctuan.com/hoc-golang/cat"
	"quoctuan.com/hoc-golang/dog"
	"quoctuan.com/hoc-golang/service"
	"quoctuan.com/hoc-golang/mouse"
)

func MakeSound(a service.Animal) {
	fmt.Printf("Day la tieng cua %s \n", a.GetName())
	fmt.Println(a.Speak())
}

func MakeSoundPlus(p service.AnimalPlus) {
	fmt.Printf("Day la tieng cua %s \n", p.GetName())
	fmt.Println(p.Speak())
	fmt.Println(p.Eat())
}

func PrintValue(val interface{}) {
	// str, ok := val.(string)
	// if ok {
	// 	fmt.Println(str)
	// } else {
	// 	fmt.Println("Please send a string")
	// }

	// numb, ok := val.(int)
	// if ok {
	// 	fmt.Println(numb)
	// } else {
	// 	fmt.Println("Please send a number")
	// }

	switch val.(type) {
	case int:
		fmt.Println(val)
	case string:
		fmt.Println(val)
	default:
		fmt.Println("Type invalid")
	}
}

func main() {
	mydog, err := dog.New("Bully")
	if err != nil {
		panic(err)
	}
	MakeSound(mydog)

	PrintValue("-=-=-=-=-=-=-=-=-=-")

	mycat, err := cat.New("Pon")
	if err != nil {
		panic(err)
	}
	MakeSoundPlus(mycat)

	PrintValue("-=-=-=-=-=-=-=-=-=-")

	mymouse, err := mouse.New("Mr Ti")
	if err != nil {
		panic(err)
	}
	PrintValue(mymouse.GetName())

	PrintValue(5)
	PrintValue(5.8)
	PrintValue(false)

}
