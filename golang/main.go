// package main

// import (
// 	"fmt"
// 	"time"
// )

// func f(from string){
// 	for i:=0;i<3;i++{
// 		fmt.Println(from,":",i)
// 	}
// }

// func main(){
// 	f("direct")

// 	go func (msg string){
// 		fmt.Println(msg)
// 	}("going")
// 	go f("goroutine")

// 	go func (msg string){
// 		fmt.Println(msg,":",msg)
// 	}("Going")

// 	time.Sleep(time.Second)
// 	fmt.Println("done")
// }

package main

import "fmt"

func main(){
	messages:=make(chan string)

	go func(){messages<-"ping"}()

	go func(){messages<-"ping2"}()
	go func(){messages<-"ping3"}()
	go func(){messages<-"ping4"}()
	go func(){messages<-"ping5"}()

	for i:=0;i<5;i++{
		msg:=<-messages


		fmt.Println(msg)
	}
}