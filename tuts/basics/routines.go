package main
import "fmt"

func f(from string) {
  for i := 0; i<3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  f("direct")

  go f("goroutine")

  go func(msg string) {
    fmt.Println(msg)
  }("going")

  var input string
  fmt.Scanln(&input) //the program will wait untill there is in input entered, then it will finish and exit
  fmt.Println("done")
}
