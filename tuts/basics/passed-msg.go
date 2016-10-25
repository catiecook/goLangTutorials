package main
import "fmt"

//this one only revieves the string message
func ping(pings chan<- string, msg string) {
  pings <- msg
}

//this one only
func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs := <-msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}
