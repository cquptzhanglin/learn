package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	var mapList map[string]int
// 	var mapCreated map[string]float32
// 	var mapAssigned map[string]int

// 	mapList = map[string]int{"one": 1, "two": 2}
// 	mapCreated = make(map[string]float32)
// 	mapAssigned = mapList

// 	mapCreated["key1"] = 4.5
// 	mapCreated["key2"] = 3.14159
// 	mapAssigned["two"] = 3

// 	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
// 	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
// 	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
// 	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])

// 	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
// 	for key := range capitals {
// 		capitals[key] = "boy"
// 		fmt.Println("Map item: Capital of", key, "is", capitals[key])
// 	}
// 	//使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
// 	var l *list.List
// 	l = list.New()
// 	l.PushBack(101)
// 	l.PushBack(102)
// 	l.PushBack(103)
// 	l.PushBack(104)
// 	//l.Back
// 	fmt.Println("Before Removing...")
// 	//遍历list，删除元素
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Println("removing", e.Value)
// 		l.Remove(e)
// 	}
// 	fmt.Println("After Removing...")
// 	//遍历删除完元素后的list
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Println(e.Value)
// 	}
// }
