package main

import "net"
import "fmt"
import "bufio"
import "strings" // for to upper case function

func main() {
    fmt.Println("## use <Ctrl+C> for exit.")
    fmt.Println(">> Waiting for new connection...")
    ln, _ := net.Listen("tcp", ":8081")
    conn, _ := ln.Accept()
    fmt.Println(">> Client connected.")
    for { //(use <Ctrl+C> for exit)
        fmt.Println(">> Waiting for client-side text input...")
        message, error := bufio.NewReader(conn).ReadString('\n')
		if error != nil {
		    conn.Close()
			break
		}
        fmt.Print(">> Message Received:", string(message))
        newmessage := strings.ToUpper(message)
        conn.Write([]byte(newmessage + "\n"))
    }
}