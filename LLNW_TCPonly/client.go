package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
    fmt.Println("## use <Ctrl+C> for exit.")
    fmt.Println(">> Dial to server...")	
    conn, error := net.Dial("tcp", "127.0.0.1:8081")
		if error != nil {
			return
		}	
    fmt.Println(">> Server connected.")
    for { 
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> Text to send: ")
        text, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, text + "\n")
        message, error := bufio.NewReader(conn).ReadString('\n')
		if error != nil {
		    conn.Close()
			break
		}
        fmt.Print(">> Message from server: "+message)
    }
}