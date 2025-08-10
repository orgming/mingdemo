package demo

const DemoKey = "demo"

type Student struct {
	ID   int
	Name string
}

type IService interface {
	GetAllStudent() []Student
}
