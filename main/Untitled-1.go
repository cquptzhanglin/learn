package main

// import (
// 	"math/rand"
// 	"time"
// )

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"math/big"
// )

// func int2bytes(num int) (b []byte) {
// 	b = make([]byte, 4)
// 	binary.BigEndian.PutUint32(b, uint32(num))
// 	return
// }
// func main() {
// 	ten := new(big.Int)
// 	ten.SetBytes(int2bytes(10))

// 	four := new(big.Int)
// 	four.SetBytes(int2bytes(4))

// 	fmt.Println(ten, four)

// 	tenmodfour := new(big.Int)
// 	tenmodfour = tenmodfour.Mod(ten, four)

// 	fmt.Println("mod", tenmodfour)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		a := rand.Int()
// 		println(a)
// 	}
// 	for i := 0; i < 5; i++ {
// 		a := rand.Intn(10)
// 		println(a)
// 	}
// 	timens := int64(time.Now().Nanosecond())
// 	rand.Seed(timens)
// 	for i := 0; i < 5; i++ {
// 		a := 100 * rand.Float32()
// 		println(a)
// 	}
// 	println(`This is a raw string \n`)
// }
