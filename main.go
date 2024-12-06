package geecache

import (
	"fmt"
	"sync"
	"time"
)

var set = make(map[int]bool)
var m sync.Mutex

func printOnce(num int) {
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}
func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
