package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sg3des/network"
)

var (
	mode string
)

func init() {
	flag.Usage = func() {
		fmt.Println(os.Args[0], " <mode>")
		fmt.Println("")
		fmt.Println("  mode:")
		fmt.Println("    server - for start server listener")
		fmt.Println("    client - for client connection")
	}
	flag.Parse()
	mode = flag.Arg(0)
	log.SetFlags(log.Lshortfile)
}

func main() {
	fmt.Println("mode:", mode)

	switch mode {
	case "server":
		server()
	case "client":
		client()
	}
}

func server() {
	_, err := network.Server(":1208") //conn,err := ...
	if err != nil {
		log.Fatal(err)
	}
	// conn.Reply(data)

	time.Sleep(10 * time.Second)
	log.Println("exit")
}

func client() {
	_, err := network.Client(":1208") //conn,err := ...
	if err != nil {
		log.Fatal(err)
	}

	p := NewPacket("client", "hi i`m client")

	// conn.Reply(username, data)
	network.Reply(p.Bytes())
	time.Sleep(5 * time.Second)
	log.Println("exit")
}

//Packet structure
type Packet struct {
	Timestamp time.Time
	Username  string
	Data      string
}

func NewPacket(username, data string) Packet {
	return Packet{
		Timestamp: time.Now().UTC(),
		Username:  username,
		Data:      data,
	}
}

//Bytes convert packet to bytes for send it - json
func (p *Packet) Bytes() []byte {
	data, _ := json.Marshal(p)
	return data
}

//RemoteCall is example of RPC - is not ready yet
func (p *Packet) RemoteCall() {
	log.Println(p)
}
