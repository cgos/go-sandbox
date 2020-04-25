package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 65138, time.UTC)
	p(then)

	secs := now.Unix()
	p(secs)

	p(now.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))

	p(e)
}
