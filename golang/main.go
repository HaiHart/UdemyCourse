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

// package main

// import "fmt"

// func main(){
// 	messages:=make(chan string)

// 	go func(){messages<-"ping"}()

// 	go func(){messages<-"ping2"}()
// 	go func(){messages<-"ping3"}()
// 	go func(){messages<-"ping4"}()
// 	go func(){messages<-"ping5"}()

// 	for i:=0;i<5;i++{
// 		msg:=<-messages

// 		fmt.Println(msg)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main(){
// 	c1:=make(chan string,1)
// 	go func ()  {
// 		time.Sleep(2*time.Second)
// 		c1<- "result 1"
// 	}()

// 	select {
// 	case res:=<-c1:
// 		fmt.Println(res)
// 	case <-time.After(1*time.Second):
// 		fmt.Println("timeout 1")
// 	}

// 	c2:=make(chan string,1)

// 	go func() {
// 		time.Sleep(2*time.Second)
// 		c2<-"result 2"
// 	}()

// 	select {
// 	case res:=<-c2:
// 		fmt.Println(res)
// 	case <-time.After(3*time.Second):
// 		fmt.Println("timeout 2")
// 	}
// }
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	messages:=make(chan string)
// 	signals:=make(chan string)

// 	select {
// 	case msg:=<-messages:
// 		fmt.Println("messages received",msg)
// 	default:
// 		fmt.Println("no messages received")

// 	}

// 	msg:="hi"

// 	select {
// 	case messages<-msg:
// 		fmt.Println("sent messages ",msg)
// 	default:
// 		fmt.Println("no messages sent")
// 	}

// 	select {
// 	case msg:=<-messages:
// 		fmt.Println("received messages ",msg)
// 	case sig:=<-signals:
// 		fmt.Println("Received signal",sig)
// 	default:
// 		fmt.Println("no activity")
// 	}
// }

// package main

// import "fmt"

// func main(){
// 	    queue := make(chan string, 2)
//     queue <- "one"
//     queue <- "two"
//     close(queue)

//     for elem := range queue {
//         fmt.Println(elem)
//     }
// }

package main

import (
	"fmt"
	"time"
)

func main(){
    ticker:=time.NewTicker(500*time.Millisecond)
    done:=make(chan bool)

    go func() {
        for{
            select {
            case <-done:  
                return  
            case t:=<-ticker.C:
                fmt.Println("Tick at: ",t)
            }
        }
    }()
    time.Sleep(1600*time.Millisecond)
    ticker.Stop()
    done<-true
    fmt.Println("Ticker stopped")
}