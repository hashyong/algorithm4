package example

import "fmt"

type programmer interface {
	WriteHelloWorld() bool
	WriteHelloWorld1() interface{}
}

func writeFirstProgram2(p programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld1())
}

func writeFirstProgram1(p programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

type i interface {
	Get() int
	Set(int)
}

type s struct {
	Age int
}

func (s s) Get() int {
	return s.Age
}

func (s *s) Set(age int) {
	s.Age = age
}

func f(i i) {
	i.Set(10)
	fmt.Println(i.Get())
}

func f1(i i) {
	i.Get()
	fmt.Println(i.Get())
}

func main() {
	s := s{}
	f(&s)  // 4
	f1(&s) // 4
}
