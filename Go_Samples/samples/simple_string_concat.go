package main

import (
    "strings"
    "fmt"
)

func stringsBuilder() {
    // ZERO-VALUE:
    //
    // It's ready to use from the get-go.
    // You don't need to initialize it.
    var sb strings.Builder
    for i := 0; i < 1000; i++ {
        sb.WriteString("a")
    }
    fmt.Println(sb.String())
}

func bytesBuffer(b *testing.B) {
    var buffer bytes.Buffer
    for n := 0; n < 1000; n++ {
        buffer.WriteString("x")
    }
}
