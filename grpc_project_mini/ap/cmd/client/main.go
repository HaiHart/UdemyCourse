package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (a Actitvity) String() string{
		return fmt.Sprintf("ID:\t\"%d\"\t%s\t%d-%d-%d",a.Id,a.Description,a.Time.Year(),a.Time.Month(),a.Time.Day())
}

func main()  {
	add:=flag.Bool("add",false,"Add activity")
	get:=flag.Bool("get",false,"Get activity")
	flag.Parse()

	switch  {
	case *get:
		id,err:=strconv.Atoi(os.Args[2])
		if err!= nil{
			fmt.Fprintln(os.Stderr,"Invalid offset: not an integer")
			os.Exit(1)
		}
		fmt.Println(id.String())

		a,err:=activitiesClient.Retrieve(id)
		if err !=nil{
			fmt.Fprintln(os.Stderr,"Error:",err.Error())
			os.Exit(1)
		}

	case *add:
		if len(os.Args)!= 3 {
			fmt.Fprintln(os.Stderr,"Usage: --add messages")
			os.Exit(1)
		}
		a:=client.Activity{
			Time:time.Now(),
			Description: os.Args[2],
		}
		id,err:=activitiesClient.Insert()

		if err != nil{
			fmt.Fprintln(os.Stderr,"Error:",err.Error())
		}

	default:
		flag.Usage()
		os.Exit(1)

	}

}