package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) new(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) log_info() {
	fmt.Println("Name : " + p.name)
	fmt.Println("Age : ", p.age)
}

func main() {
	p := Person{}
	p.new("Joe", 20)
	p.log_info()
}
