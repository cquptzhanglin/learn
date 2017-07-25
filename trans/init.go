package trans

import "math"

//Pi gloabl pi
var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

/*
func int2bytes(num int) (b []byte) {
    b = make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(num))
    return
}
func main() {
    ten := new(big.Int)
    ten.SetBytes(int2bytes(10))

    four := new(big.Int)
    four.SetBytes(int2bytes(4))

    fmt.Println(ten, four)

    tenmodfour := new(big.Int)
    tenmodfour = tenmodfour.Mod(ten, four)

    fmt.Println("mod", tenmodfour)
}
*/
