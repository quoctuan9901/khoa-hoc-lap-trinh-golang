package service

type Animal interface {
	Speak() string
	GetName() string
}

type AnimalPlus interface {
	Animal
	Eat() string
}

type AnimalAction interface {
	Run() string
}