package main

import(
	//"fmt"
	//"io/ioutil"
	"net"
	//"os"
	"log"
	"time"
)

/*func main(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}*/

func establishConn(i int) net.Conn {
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        log.Printf("%d: dial error: %s", i, err)
        return nil
    }
    log.Println(i, ":connect to server ok")
    return conn
}

func main() {
    var sl []net.Conn
    for i := 1; i < 1000; i++ {
        conn := establishConn(i)
        if conn != nil {
            sl = append(sl, conn)
        }
    }

    time.Sleep(time.Second * 10000)
}