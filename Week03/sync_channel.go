package main

import(
    "fmt"
)




var ch = make(chan int, 10)

var a string

func f(){
    a = "hello, world"
    ch <- 0
}

func main(){
    go f()
    fmt.Println(a)
}
