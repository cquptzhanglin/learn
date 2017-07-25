package main

// import (
// 	"fmt"
// 	"time"
// )

// var week time.Duration

// func main() {
// 	t := time.Now()
// 	fmt.Println(t)
// 	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

// 	t = time.Now().UTC()
// 	fmt.Println(t)
// 	fmt.Println(time.Now())

// 	fmt.Println("newline")
// 	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
// 	week_from_now := t.Add(week)
// 	fmt.Println("duaration", week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
// 	// formatting times:
// 	fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
// 	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
// 	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
// 	s := t.Format("20060102")
// 	fmt.Println(t, "=>", s)

// 	var i int
// 	var intP, ptr *int
// 	i = 5
// 	intP = &i
// 	fmt.Printf("%d %p %d %p\n", i, intP, *intP, ptr)
// }
