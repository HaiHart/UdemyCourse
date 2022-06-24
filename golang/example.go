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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main(){
//     ticker:=time.NewTicker(500*time.Millisecond)
//     done:=make(chan bool)

//     go func() {
//         for{
//             select {
//             case <-done:
//                 return
//             case t:=<-ticker.C:
//                 fmt.Println("Tick at: ",t)
//             }
//         }
//     }()
//     time.Sleep(1600*time.Millisecond)
//     ticker.Stop()
//     done<-true
//     fmt.Println("Ticker stopped")
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func worker(id int, jobs <- chan int, results chan <- int){
//     for j:=range jobs{
//         fmt.Println("Worker ",id," started job ",j)
//         time.Sleep(time.Second)
//         fmt.Println("Worker ",id," finished jobs ",j)
//         results<-j*2;
//     }
// }

// func main(){
//     start :=time.Now().Second()
//     const numJobs=5
//     jobs:=make(chan int,numJobs)
//     results:=make(chan int,numJobs)

//     for w:=1;w<=3;w++{
//         go worker(w,jobs,results)
//     }

//     for j:=1;j<=numJobs;j++{
//         jobs<-j
//     }
//     close(jobs)

//     for a:=1;a<=numJobs;a++{
//         <-results
//     }
//     fmt.Println(time.Now().Second()-start)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int){
//     fmt.Println("Worker %d starting\n",id)
//     time.Sleep(time.Second)

//     fmt.Println("Worker %d finished",id)
// }

// func main()  {
//     var start=time.Now().Second()
//     var wg sync.WaitGroup

//     for i:=1;i<=5;i++{
//         wg.Add(1)

//         i:=i

//         go func() {
//             defer wg.Done()
//             worker(i)
//         }()
//     }
//     wg.Wait()
//     fmt.Println(time.Now().Second()-start)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main(){

//     var start = time.Now().Second()

//     requests:=make(chan int,5)

//     for i := 1; i <= 5; i++ {
//         requests<-i
//     }
//     close(requests)

//     limiter:=time.Tick(200*time.Millisecond)

//     for req:=range requests{
//         <-limiter
//         fmt.Println("Request ",req)
//     }

//     burstyLimiter:= make(chan time.Time,3)

//     for i := 1; i <= 3; i++ {
//         burstyLimiter<-time.Now()
//     }

//     go func() {
//         for t := range time.Tick(200*time.Microsecond) {
//             burstyLimiter<-t
//         }
//     }()

//     burstyRequests:=make(chan int, 5)

//     for i := 1; i < 6; i++ {
//         burstyRequests<-i
//     }

//         close(burstyRequests)

//     for req:=range burstyRequests{
//         <-burstyLimiter
//         fmt.Println("request ", req, time.Now())
//     }

//     fmt.Println(time.Now().Second()-start)
// }

// package main

// import (
//     "fmt"
//     "sync"
//     "sync/atomic"
// )

// func main() {

//     var ops uint64

//     var wg sync.WaitGroup

//     for i := 0; i < 50; i++ {
//         wg.Add(1)

//         go func() {
//             for c := 0; c < 1000; c++ {

//                 atomic.AddUint64(&ops, 1)
//                 // ops++
//             }
//             wg.Done()
//         }()
//     }

//     wg.Wait()

//     fmt.Println("ops:", ops)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Container struct{
//     mu  sync.Mutex
//     counters    map[string]int
// }

// func (c *Container) inc (name string){
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     c.counters[name]++
// }

// func main(){
//     c:= Container{

//         counters: map[string]int{"a": 0, "b": 0},
//     }

//     var wg sync.WaitGroup

//     doIncrement:= func (name string,n int)  {
//         for i := 0; i < n; i++ {
//             c.inc(name )
//         }
//         wg.Done()
//     }
//     wg.Add(4)
//     go doIncrement("a",5)
//     go doIncrement("b",6)
//     go doIncrement("a",7)
//     go doIncrement("c",7)
//     wg.Wait()
//     fmt.Println(c.counters)
// }

// package main

// import (
// 	"math/rand"
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// type readOp struct{

//     key int
//     resp chan int
// }

// type writeOp struct{
//     key int
//     val int
//     resp chan bool
// }

// func main (){
//     var readOps uint64
//     var writeOps uint64
//     reads:=make(chan readOp)
//     writes:= make(chan writeOp)

//     go func() {
//         var state = make(map[int]int)

//         for{
//             select {
//             case read:=<-reads:
//                 read.resp<-state[read.key]
//             case write:=<-writes:
//                 state[write.key]=write.val
//                 write.resp<-true
//             }

//         }
//     }()

//     for r := 0; r < 100; r++ {
//         go func() {
//             for{
//                 read:=readOp{
//                     key: rand.Intn(5),
//                     resp: make(chan int),
//                 }
//                 reads<-read
//                 <-read.resp
//                 atomic.AddUint64(&readOps,1)
//                 time.Sleep(time.Millisecond)

//             }
//         }()
//     }

//     for w:=0;w<10;w++{
//         go func() {
//             for {
//                 write:=writeOp{
//                     key: rand.Intn(5),
//                     val: rand.Intn(100),
//                     resp: make(chan bool),
//                 }
//                 writes<-write
//                 <-write.resp
//                 atomic.AddUint64(&writeOps,1)
//                 time.Sleep(time.Millisecond)
//             }
//         }()
//     }
//     time.Sleep(time.Second*2)

//     readOpsFinal:=atomic.LoadUint64(&readOps)
//     fmt.Println("readOps",readOpsFinal)
//     writeOpsFinal:=atomic.LoadUint64(&writeOps)
//     fmt.Println("writeOps",writeOpsFinal)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main(){
//     strs:=[]string{"c","a","b"}

//     sort.Strings(strs)
//     fmt.Println("strings:",strs)

//     ints:=[]int{7,4,3,5,2}
//     sort.Ints(ints)
//     fmt.Println("ints:",ints)

//     s:=sort.IntsAreSorted(ints)
//     fmt.Println("s:",s)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type byLength []string

// func (b byLength) Len ()int{
//     return len(b)
// }

// func (b byLength) Swap (i,j int){
//     b[i],b[j]=b[j],b[i]
// }

// func (b byLength) Less (i,j int) bool{
//     return len(b[i])<len(b[j])
// }

// func main(){
//     fruits:=[]string {"peach","banana","mango"}
//     sort.Sort(byLength(fruits))
//     fmt.Println(fruits)
// }
// package main

// import (
//     "os"
//     "text/template"
// )

// func main() {

//     t1 := template.New("t1")
//     t1, err := t1.Parse("Value is {{.}}\n")
//     if err != nil {
//         panic(err)
//     }

//     t1 = template.Must(t1.Parse("Value: {{.}}\n"))

//     t1.Execute(os.Stdout, "some text")
//     t1.Execute(os.Stdout, 5)
//     t1.Execute(os.Stdout, []string{
//         "Go",
//         "Rust",
//         "C++",
//         "C#",
//     })

//     Create := func(name, t string) *template.Template {
//         return template.Must(template.New(name).Parse(t))
//     }

//     t2 := Create("t2", "Name: {{.Name}}\n")

//     t2.Execute(os.Stdout, struct {
//         Name string
//     }{"Jane Doe"})

//     t2.Execute(os.Stdout, map[string]string{
//         "Name": "Mickey Mouse",
//     })

//     t3 := Create("t3",
//         "{{if . -}} {{ . }} {{else -}} no ! {{end}}\n")
//     t3.Execute(os.Stdout, "not empty")
//     t3.Execute(os.Stdout, "")

//     t4 := Create("t4",
//         "Range: {{range .}}{{.}} {{end}}\n")
//     t4.Execute(os.Stdout,
//         []string{
//             "Go",
//             "Rust",
//             "C++",
//             "C#",
//         })
// }
var y sync.WaitGroup

y.Add(2)

y.Done()

y.Wait()