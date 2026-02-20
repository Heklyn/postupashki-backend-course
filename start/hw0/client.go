package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("failed to connect, err:", err)
        os.Exit(1)
    }
    defer conn.Close()

    correct_response := "OK\n"

    reader := bufio.NewReader(conn)

    response, err := reader.ReadBytes(byte('\n'))
    
    if err != nil {
        fmt.Println("failed to read response, err:", err)
        os.Exit(1)
    }
    
    if string(response) != correct_response {
        fmt.Printf("incorrect response: %q, (need %q)\n", response, correct_response)
        os.Exit(1)
    }

    fmt.Printf("get correct response: %q\n", response)
    fmt.Println("close connection")
}
