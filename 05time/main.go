package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	then := time.Now()
	fmt.Println("Then", then.Format("01-02-2006 15:04:05 PM Monday"))

	time.Sleep(3 * time.Second)

	now := time.Now()
	fmt.Println("Now", now.Format("01-02-2006 03:04:05 PM Mon"))

	elapsed := now.Sub(then)
	fmt.Println("Elapsed", elapsed)

	fmt.Println("Since", time.Since(then))

	createdAt := time.Date(2006, time.Month(1), 2, 13, 34, 12, 312, time.Local)
	fmt.Println("Created At", createdAt.Format("01-02-2006 03-04-05 PM Monday"))
}
