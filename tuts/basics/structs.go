package main
import "fmt"

type person struct {
  name string
  age int
}

func main() {
  fmt.Println(person{"bob", 20})
  fmt.Println(person{name: "alice", age: 30})
  fmt.Println(person{name: "fred"})
  fmt.Println(person{name: "sean", age: 50})

  s:= person{name: "Anne", age: 50}
  fmt.Println(s.name)
  sp := &s
  fmt.Println(s.name)

  sp.age = 51
  fmt.Println(sp.age)
}
