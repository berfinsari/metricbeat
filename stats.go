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
	var atama [5]string
	var kontrol [5]string

	conn, err := net.Dial("tcp",hostname+":"+portnum)
	if err != nil{
		fmt.Println("Error %v",err)
		return
	} else{
		defer conn.Close()
		fmt.Fprintf(conn,command+"\n")
		message,_ = bufio.NewReader(conn).ReadString('E')
	}
	veri := strings.SplitAfter(message, "\n")
	//çıktıları satırlara göre ayırdık.
	for i:=0;i<len(veri)-1;i++ {
		gecici := strings.Fields(veri[i])
		//satırları boşluklara göre ayırır. ilki STAT 2. ismi 3. değeri
		atama[i] = gecici[2]
		kontrol[i] = gecici[1]
	}
	fmt.Println(atama[1], kontrol[1])
	event := common.Mapstr{
		"active_slabs" : 
		"total_malloced" : 
}
