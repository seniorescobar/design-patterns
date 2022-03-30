package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/singleton/logger"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			l := logger.GetLogger()
			l.Log("Hello, World!")
		}()
	}

	fmt.Scanln()
}
