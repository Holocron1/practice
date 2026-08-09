package main

import "fmt"

// Person содержит имя и возраст человека.
type Person struct {
	Name string
	Age  int
}

// TODO: реализуйте Stringer для Person

func main() {
	alice := Person{Name: "Alice", Age: 30}
	fmt.Println(alice)
	// Ожидаемый вывод: Alice (30)
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age) //пришлось узнать как работает Sprintf
}

type Stringer interface {
	String() string
}
