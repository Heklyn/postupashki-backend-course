package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    port := ":8080"

    listener, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Println("failed to create listener, err:", err)
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Printf("listening on %s\n", listener.Addr())

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("failed to accept connection, err:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    _, err := conn.Write([]byte("OK\n"))

    if err != nil {
        fmt.Println("failed to response, err:", err)
        return
    }
}


