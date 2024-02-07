package main
import (
    "encoding/json"
    "fmt"
    "os"
    "net"
	"io"
	"strconv"
)
// User is test for json
type configuration struct {
    Port int
    Ip    string
    Listen int
}
func tcp_handle(tcpConn net.Conn,dst_addr string){
	remote_tcp,err:=net.Dial("tcp",dst_addr) //连接目标服务器
	if err!=nil{
		fmt.Println(err)
		return
	}
	go io.Copy(remote_tcp,tcpConn) 
	go io.Copy(tcpConn,remote_tcp)
}

func tcp_listen(local_addr string,dst_addr string){
	ln,err:=net.Listen("tcp",local_addr)
	if err!=nil{
		fmt.Println("tcp_listen:",err)
		return
	}
	defer ln.Close()
	for {
		tcp_Conn,err:=ln.Accept() //接受tcp客户端连接，并返回新的套接字进行通信
		if err!=nil{
			fmt.Println("Accept:",err)
			return
		}
		go tcp_handle(tcp_Conn,dst_addr)   //创建新的协程进行转发
	}
}


func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
    if err != nil {
        fmt.Println("Error:", err)
    }
	fmt.Println("configuration example:{\"listen\": 58060,\"ip\": \"47.102.129.85\",\"port\": 9527}")
    fmt.Println("ip:" , conf.Ip)
    fmt.Println("Listen:" , conf.Listen)
    fmt.Println("Port:" , conf.Port)

    fmt.Println("this is porxy server:",conf.Listen)
	go tcp_listen("0.0.0.0:" + strconv.Itoa(conf.Listen),conf.Ip + ":" + strconv.Itoa(conf.Port))
	for{}
}