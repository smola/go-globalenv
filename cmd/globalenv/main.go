package main

import (
	"fmt"

	"github.com/smola/go-globalenv"
)

func main() {
	for _, e := range globalenv.Environ() {
		fmt.Println(e)
	}
}
