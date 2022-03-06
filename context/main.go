package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	exampleWithValue(ctx)
}

func exampleWithValue(ctx context.Context) {
	valueCtx := context.WithValue(ctx, "key1", "value1")
	fmt.Println(valueCtx.Value("key1").(string))
}
