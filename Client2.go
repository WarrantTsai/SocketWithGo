package main

import(
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	"log"
	"time"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("usage: go run client2.go YOUR_CONTENT")
        return
    }
    log.Println("begin dial...")
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")
	data := os.Args[1]
	for true{
		time.Sleep(time.Second * 2)
		conn.Write([]byte(data))
		fmt.Printf("Say something: ")
		fmt.Scan(&data)
	}
}