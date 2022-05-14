package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	fmt.Println("Hello World!")
	cancel()

	<-ctx.Done()
}
