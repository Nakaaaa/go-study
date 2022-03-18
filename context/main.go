package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	ctx := context.Background()
// 	exampleWithValue(ctx)
// }

// func exampleWithValue(ctx context.Context) {
// 	valueCtx := context.WithValue(ctx, "key1", "value1")
// 	fmt.Println(valueCtx.Value("key1").(string))
// }

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)
	// cancel()

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("######################")
}
