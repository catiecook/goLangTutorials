package main
import "fmt"

func main() {
  //channels are typed by the values they convey.
  //make a new channel with this syntax

  messages := make(chan string)

//this is a function that puts the values ping into messages
// so now that messages is a new channel, it can recieve the value "ping" with the arrow pointer <-
  go func() { messages <- "ping" } ()

//setting the vaiable msg to equal messages which is holding the string value "ping"
//we have to set the value this way because messages (the channel) is pointing to the address, so if we printed it, it would print the address of ping, nto the value "ping"
  msg := <-messages

  fmt.Println(msg)
}
