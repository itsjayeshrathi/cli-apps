package main

import (
	"context"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type key string

func someFunc(ctx context.Context) context.Context {
	return context.WithValue(ctx, key("user"), User{
		Name: "John Doe",
		Age:  20,
	})
}

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	ctx = context.TODO()
	fmt.Println(ctx)
}
