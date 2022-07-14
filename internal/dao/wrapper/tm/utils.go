package tm

import (
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

func RandomInt(n int) int64 {
	var l = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var v string
	rand.Seed(time.Now().UnixNano())
	b := make([]int64, n)
	length := len(l)
	for range b {
		v = v + cast.ToString(l[rand.Intn(length)])
	}
	return cast.ToInt64(v)
}
