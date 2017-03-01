package main

import "testing"
import "fmt"
import _ "net/http/pprof"

func test(t *testing.T) {
	fmt.Print("Hello, world test")
}
