package main

import (
    "bufio"
    "fmt"
    "net"
)

func main(){
    command := "stats"
    hostName := "localhost"
    portNum := "11211"

    doDial(command, hostName, portNum)
    //fmt.Println(deneme)
}
func doDial(cmd, host, port string) {
    conn, err := net.Dial("tcp", host+":"+port)
    if err != nil {
        fmt.Printf("Error %v", err)
        return
    } else {
        defer conn.Close()
        fmt.Fprintf(conn,cmd+"\n")
        message,_:=bufio.NewReader(conn).ReadString('E')
        fmt.Print("\n" + message)
    }
}
