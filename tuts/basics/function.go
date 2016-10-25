package main
import "fmt"

// one return value

func plus(a int, b int) int {
  return a + b
}

func plusPlus(a, b, c int) int {
  return a + b + c
}

//multip return values

func vals() (int, int) {
  return 3, 7
}

func main() {
  res := plus(1, 2)
  fmt.Println("1+2 =", res)

  res = plusPlus(1, 2, 3)
  fmt.Println("1+2+3 =", res)

  one, two := vals()
  fmt.Println(one)
  fmt.Println(two)

  _, three := vals()
  fmt.Println(three)
}
