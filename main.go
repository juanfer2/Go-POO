package main

import "fmt"

type Employee struct {
	Id int
	Name string
}

// Constructor
func NewEmployee(id int, name string) *Employee{
	return &Employee{
		Id: id, 
		Name: name,
	}
}

// Behaviors
func(e *Employee) GetId() int {
	return e.Id
} 

func(e *Employee) GetName() string {
	return e.Name
} 

func(e *Employee) SetId(id int) {
	e.Id = id
} 

func(e *Employee) SetName(name string) {
	e.Name = name
}

func(e Employee) GetMessage() string {
	return "Full time employee"
}

// Herency
type TemporalEmployee struct {
	Employee
}

func(te TemporalEmployee) GetMessage() string {
	return "temporal time"
}

// Polymorphis
type getMessageInterface interface {
	GetMessage() string
}

func GetMessage(g getMessageInterface) string {
	return g.GetMessage()
}

func main() {
	employee := Employee{}

	fmt.Println(employee.Id)
	fmt.Println(employee.Name)

	employee.SetId(1)
	employee.SetName("Jhon Doe")

	fmt.Println(employee.Id)
	fmt.Println(employee.Name)

	fmt.Println(employee.GetId())
	fmt.Println(employee.GetName())

	emp := NewEmployee(1, "Juanfer") 
	fmt.Println(emp.GetName())
	fmt.Println(emp.GetId())

	temp := TemporalEmployee{}

	temp.SetId(3)
	temp.SetName("Temporal name")
	fmt.Println(temp.GetName())
	fmt.Println(temp.GetId())

	fmt.Println(GetMessage(temp))
	fmt.Println(GetMessage(emp))
}
