package main

import(
	"bufio"
	"fmt"
	"net"
	"strings" 
)
func main() {
	command := "stats slabs"
	hostname := "localhost"
	portnum := "11211"
	var message string

	conn, err := net.Dial("tcp",hostname+":"+portnum)
	if err != nil{
		fmt.Println("Error %v",err)
		return
	} else{
		defer conn.Close()
		fmt.Fprintf(conn,command+"\n")
		message,_ = bufio.NewReader(conn).ReadString('E')
	}
//	fmt.Println(message)
	veri := strings.SplitAfter(message, "\n")
	fmt.Println(veri[1])
	for i:=0;i<len(veri)-1;i++ {
		gecici := strings.Fields(veri[i])
		fmt.Println(gecici[2])
	}

}
