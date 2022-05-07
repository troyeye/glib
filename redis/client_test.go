package gredis

import (
	"fmt"
	"testing"
	"time"
)

func Test_getClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(url int) {
			defer RecoverHandle(url)
			getClient(fmt.Sprintf("%d", url))
		}(i)
	}

	go func() {
		for {
			fmt.Println(len(clients))
			time.Sleep(time.Second)
		}

	}()
	time.Sleep(time.Second*5)
}

func RecoverHandle(v ...interface{}) {
	if err := recover(); err != nil {
		fmt.Println("111111111111111111111111111111111111111111")
	}
}
