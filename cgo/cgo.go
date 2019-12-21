package cgo

// int sum(int a, int b) {
//     return a + b;
// }
import "C"

// Sum C 语言实现的加法运算
func Sum(a, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
