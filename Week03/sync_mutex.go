package main

import(
    "fmt"
    "sync"
)


// // bad case: out put ""
// var s string
//
// func f() {
//     s = "hello world"
// }
//
// func main() {
//     go f()
//     fmt.Println(s)
// }


// good case, output: hello world
var l sync.Mutex
var s string

func f() {
    s = "hello world"
    l.Unlock()
}

func main() {
    l.Lock()
    go f()
    l.Lock()
    fmt.Println(s)
}
