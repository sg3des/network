package network

import (
	"log"
	"net"
)

var c *Connection

//udpAddr return resolved addr for udp connection
func udpAddr(addr string) *net.UDPAddr {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	return raddr
}

//Server - start server listener and return connection
func Server(addr string) (*Connection, error) {
	conn, err := net.ListenUDP("udp", udpAddr(addr))
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	c = &Connection{
		Conn: conn,
	}

	go c.carrier()
	// c =
	return c, nil
}

//Client - initialize connection and return connection
func Client(addr string) (*Connection, error) {
	conn, err := net.DialUDP("udp", udpAddr(":0"), udpAddr(addr))
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	c = &Connection{
		Conn: conn,
	}

	go c.carrier()
	return c, nil
}

//carrier manage incoming messages
func (c *Connection) carrier() {
	var b = make([]byte, 2048)
	for {
		n, addr, err := c.Conn.ReadFromUDP(b)
		if err != nil {
			log.Println("read from the closed connection", err)
			return
		}
		b = append([]byte(addr.String()+" "), b[:n]...)
		log.Println(string(b))
	}
}

//Close global connection
func Close() {
	if c == nil {
		return
	}
	c.Close()
}

//Close connection
func (c *Connection) Close() {
	if c.Conn == nil {
		return
	}
	err := c.Conn.Close()
	if err != nil {
		log.Println("failed close connection", err)
	}
}

//Connection structure
type Connection struct {
	Conn *net.UDPConn
}

//Reply - send message from global connection
func Reply(data []byte) {
	if c == nil {
		log.Fatal("connection is not established")
	}
	c.Reply(data)
}

//Reply - send message
func (c *Connection) Reply(data []byte) {
	_, err := c.Conn.Write(data)
	if err != nil {
		log.Println("write to closed connection", err)
	}
}
