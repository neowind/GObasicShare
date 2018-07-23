package routine

import (
	"fmt"
	"time"
)

func Count(key string, sleep int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("%s: [%d]\n", key, i)
		time.Sleep(time.Second * time.Duration(sleep))
	}
}
