package main
import "fmt"

//arguments will be passed by value into this function, then ival will be holding a copy of the ival int distinct from the one in the calling function

//DOES NOT CHANGE VALUE OF I
func zeroval(ival int) {
  ival = 0
}

//this funciton takes in an integer pointer. The * differences the pointer from its memory address to the current value at that address.
//**Assignming a value to a dereferenced pointer changes the value at the referenced address**

//CAN CHANGE VALUE OF I BECAUSE IT HAS THE ORIGINAL ADDRESS TO IT
func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)
//The i will make the variable hold the address of the i, not the actual value.
  fmt.Println("pointer:", &i)
}
