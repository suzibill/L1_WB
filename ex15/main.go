package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = string(v[:100])
	tmp := make([]byte, 0, 100)
	copy(tmp, v)
	justString = string(tmp)

}

func createHugeString(size int) []byte {
	return make([]byte, 0, size)
}

func main() {
	someFunc()
}

// проблема в том что мы создаём большой слайс, после его перенарезаем
// и получается что у нас большое лишнее капасити, которое мы не используем
