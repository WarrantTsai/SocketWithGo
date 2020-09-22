package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"log"
)

func main(){
	/*port := ":8888"
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	for{
		conn, err :=listener.Accept()
		if err != nil{
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}*/
	
	l, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Println("error listen:", err)
        return
    }
    defer l.Close()
    log.Println("listen ok")

    var i int
    for {
        time.Sleep(time.Second * 10)
        if _, err := l.Accept(); err != nil {
            log.Println("accept error:", err)
            break
        }
        i++
        log.Printf("%d: accept a new connection\n", i)
	}
}

func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}