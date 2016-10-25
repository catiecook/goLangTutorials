package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  f := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World!")
  }

  http.Fprintf("/", f)

  log.Panic(http.ListenAndServe(":3000", nil))
}
