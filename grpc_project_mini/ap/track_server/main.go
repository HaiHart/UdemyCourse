package track_server

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	. "github.com/HaiHart/UdemyCourse/tree/main/grpc_project_mini/ap/api/api"
	"github.com/gorilla/mux"
)



// type Activity struct {
// 	time time.Time	`json:"time"`
// 	Description string	`json:"Description"`
// 	Id uint64	`json:"id"`
// }

// type Activities struct {
// 	activities []Activity
// }

// type IDDocument struct{
// 	ID uint64 `json:"id"`
// }

// type ActivityDocument struct{
// 	activity Activity `json:"Activity"`
// }

type httpServer struct{
	activities *Activities
}

func (c *Activities)Insert(activity Activity) uint64 {
	activity.Id=uint64(len(c.activities))
	c.activities=append(c.activities, activity)	
	return activity.Id
}

func (c *Activities) Retrieve(id uint64) (Activity,error) {
	var ErrIDNotFound= fmt.Errorf("Id not found")	
	if id>=uint64(len(c.activities)){
		return Activity{},ErrIDNotFound
	}
	return c.activities[id],nil
}

func (s *httpServer)handleGet(w http.ResponseWriter,req *http.Request)  {
	fmt.Fprintf(w,"get\n")
}

func (s*httpServer) handlePost(w http.ResponseWriter, r *http.Request)  {
	var req ActivityDocument
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(	w, err.Error(),http.StatusBadRequest)
		return
	}
	id:=s.activities.Insert(req.activity)
	res:=IDDocument{
		ID: id,
	}
	json.NewEncoder(w).Encode(res)
}

func NewHTTPServer(addr string) *http.Server {
	server:=&httpServer{
		activities:&Activities{},
	}
	r:=mux.NewRouter()
	r.HandleFunc("/",server.handleGet).Methods("GET")
	r.HandleFunc("/",server.handlePost).Methods("POST")
	return &http.Server{
		Addr: addr,
		Handler: r,
	}
}

func main()  {
	add:=flag.Bool("add",false,"Add activity")
	get:=flag.Bool("get",false,"Get activity")
	flag.Parse()

	switch  {
	case *get:
		
	case *add:
		if len(os.Args)!= 3 {
			fmt.Fprintln(os.Stderr,"Usage: --add messages")
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)

	}

	srv:=NewHTTPServer(":8000")
	srv.ListenAndServe()
}