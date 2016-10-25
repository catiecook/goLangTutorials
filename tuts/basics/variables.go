package main
import "fmt"

func main() {
  var a string = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b,c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f:="short"
  // := is shorthand for declaring and initializing a variable
  //so it tells us that f is a string that equals short here
  
  fmt.Println(f)
}
