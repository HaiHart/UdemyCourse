package api

import "time"

type Activity struct {
	Time time.Time	`json:"time"`
	Description string	`json:"Description"`
	Id uint64	`json:"id"`
}

type Activities struct {
	activities []Activity
}

type IDDocument struct{
	ID uint64 `json:"id"`
}

type ActivityDocument struct{
	Activity Activity `json:"Activity"`
}


