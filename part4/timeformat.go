package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {

	now := time.Now()
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if e != nil {
		panic(e)
	}
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))

	p("")
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
}
