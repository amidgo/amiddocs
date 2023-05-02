package amidstr

import "math/rand"

func MakeString(size int) string {
	if size == 0 {
		return ""
	}
	symbols := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, size)
	for i := range b {
		rint := rand.Intn(len(symbols))
		b[i] = symbols[rint]
	}
	return string(b)
}
