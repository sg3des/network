package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net"
	"reflect"
	"regexp"
	"strings"
)

var (
	c *Connection

	rpc = make(map[string]RPC)
)

//udpAddr return resolved addr for udp connection
func udpAddr(addr string) *net.UDPAddr {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalln("failed addr", err)
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
	var bRequest = make([]byte, 2048)
	for {
		n, addr, err := c.Conn.ReadFromUDP(bRequest)
		if err != nil {
			log.Println("read from the closed connection", err)
			return
		}
		request, err := parseRequest(bRequest[:n])
		if err != nil {
			log.Println(err)
			continue
		}

		rpcStruct, method := request.getRPC()
		if rpcStruct == nil || !method.IsValid() || method.IsNil() { //
			log.Println("request not valid", addr, request.Struct, request.Method)
			continue
		}

		request.call(rpcStruct, method)
	}
}

//request incoming data structure
type request struct {
	Struct string
	Method string
	Packet []byte
}

//parseRequest return parse request from []byte incoming data
func parseRequest(bRequest []byte) (request, error) {
	if len(bRequest) == 0 {
		return request{}, errors.New("length of request 0")
	}
	splitted := bytes.SplitN(bRequest, []byte(" "), 2)
	if len(splitted) != 2 {
		return request{}, errors.New("request is not valid")
	}

	rpcname := bytes.SplitN(splitted[0], []byte("."), 2)
	if len(rpcname) != 2 {
		return request{}, errors.New("rpcname in request not valid")
	}

	r := request{
		Struct: strings.Trim(string(rpcname[0]), " \t."),
		Method: strings.Trim(string(rpcname[1]), " \t."),
		Packet: splitted[1],
	}

	return r, nil
}

//getRPC check request on valid and store in global map new methods
func (r *request) getRPC() (interface{}, reflect.Value) {
	rpc, ok := rpc[r.Struct]
	if !ok {
		log.Println("rpc structure not exist")
		return nil, reflect.Value{}
	}

	if method, ok := rpc.Methods[r.Method]; ok {
		return rpc.Interface, method
	}

	method := rpc.Struct.MethodByName(r.Method)
	if !method.IsValid() {
		log.Println("rpc method not valid")
		return nil, reflect.Value{}
	}
	rpc.Methods[r.Method] = method

	return rpc.Interface, method
}

//call function
func (r *request) call(rpcInterface interface{}, method reflect.Value) {
	if len(r.Packet) > 1 {
		err := json.Unmarshal(r.Packet, &rpcInterface)
		if err != nil {
			log.Println("failed unmarshal data", string(r.Packet), "reason:", err)
		}
	}

	method.Call([]reflect.Value{})
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
func Reply(rpcname string, data []byte) {
	if c == nil {
		log.Fatal("connection is not established")
	}
	c.Reply(rpcname, data)
}

//Reply - send message
func (c *Connection) Reply(rpcname string, data []byte) {
	data = append([]byte(rpcname+" "), data...)
	_, err := c.Conn.Write(data)
	if err != nil {
		log.Println("write to closed connection", err)
	}
}

//SetRPC bind rpc functions
func SetRPC(userRPC ...interface{}) {
	for _, c := range userRPC {
		rpc[getRPCname(c)] = newRPC(c)
	}
}

//RPC struct contains prepared for call reflect of functions
type RPC struct {
	Interface interface{}
	Struct    reflect.Value
	Methods   map[string]reflect.Value
}

func newRPC(c interface{}) (rpc RPC) {
	rpc.Interface = c
	rpc.Struct = reflect.ValueOf(c)
	rpc.Methods = make(map[string]reflect.Value)

	return
}

//getRPCname determine name of struct through reflect
func getRPCname(c interface{}) string {
	constring := reflect.ValueOf(c).String()
	f := regexp.MustCompile("\\.([A-Za-z0-9]+) ").FindString(constring)
	return strings.Trim(f, ". ")
}
